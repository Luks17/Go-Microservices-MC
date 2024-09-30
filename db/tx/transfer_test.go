package tx_test

import (
	"context"
	"testing"

	"github.com/Luks17/Go-Microservices-MC/db/devutils"
	"github.com/Luks17/Go-Microservices-MC/db/tx"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := tx.NewStore(testDB)

	account1 := devutils.RandomNewAccount()
	account2 := devutils.RandomNewAccount()

	createdAccount1, err := testQueries.CreateAccount(context.Background(), account1)

	require.NoError(t, err)
	require.NotEmpty(t, createdAccount1)

	createdAccount2, err := testQueries.CreateAccount(context.Background(), account2)

	require.NoError(t, err)
	require.NotEmpty(t, createdAccount2)

	// number of transfers
	n := 5
	// amount to be transferred
	amount := "10.50"

	errs := make(chan error)
	results := make(chan tx.TransferTxResult)

	// goroutines for transfers
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), tx.TransferTxParams{
				FromAccountID: createdAccount1.ID,
				ToAccountID:   createdAccount2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// checking results from goroutines
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer params
		require.NotEmpty(t, result.Transfer)
		require.NotZero(t, result.Transfer.ID)

		transfer, err := store.GetTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, err)
		require.NotEmpty(t, transfer)

		require.Equal(t, createdAccount1.ID, transfer.FromAccountID)
		require.Equal(t, createdAccount2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		// check fromEntry params
		require.NotEmpty(t, result.FromEntry)
		require.NotZero(t, result.FromEntry.ID)

		fromEntry, err := store.GetEntry(context.Background(), result.FromEntry.ID)
		require.NoError(t, err)
		require.NotEmpty(t, fromEntry)
		require.Equal(t, createdAccount1.ID, fromEntry.AccountID)
		require.Equal(t, "-"+amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		// check toEntry params
		require.NotEmpty(t, result.ToEntry)
		require.NotZero(t, result.ToEntry.ID)

		toEntry, err := store.GetEntry(context.Background(), result.ToEntry.ID)
		require.NoError(t, err)
		require.NotEmpty(t, toEntry)
		require.Equal(t, createdAccount2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)
	}
}
