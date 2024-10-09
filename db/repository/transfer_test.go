package repository_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/Luks17/Go-Microservices-MC/db/repository"
	"github.com/Luks17/Go-Microservices-MC/db/sqlc"
	"github.com/Luks17/Go-Microservices-MC/devutils"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := repository.NewStore(testDB)

	createdAccount1 := newAccount(t)
	createdAccount2 := newAccount(t)

	oldBalance1, err := strconv.ParseFloat(createdAccount1.Balance, 64)
	require.NoError(t, err)
	oldBalance2, err := strconv.ParseFloat(createdAccount2.Balance, 64)
	require.NoError(t, err)

	// number of transfers
	n := 5
	// amount to be transferred
	amount := "10.50"

	errs := make(chan error)
	results := make(chan repository.TransferTxResult)

	// goroutines for transfers
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), repository.TransferTxParams{
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
		require.Equal(t, createdAccount1.ID, fromEntry.AccountID)
		require.Equal(t, "-"+amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		// check toEntry params
		require.NotEmpty(t, result.ToEntry)
		require.NotZero(t, result.ToEntry.ID)

		toEntry, err := store.GetEntry(context.Background(), result.ToEntry.ID)
		require.NoError(t, err)
		require.Equal(t, createdAccount2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		// check fromAccount Params
		require.NotEmpty(t, result.FromAccount)
		require.NotZero(t, result.ToAccount.ID)
		require.Equal(t, createdAccount1.ID, result.FromAccount.ID)

		// check toAccount Params
		require.NotEmpty(t, result.ToAccount)
		require.NotZero(t, result.ToAccount.ID)
		require.Equal(t, createdAccount2.ID, result.ToAccount.ID)

		// check balances
		newBalance1, err := strconv.ParseFloat(result.FromAccount.Balance, 64)
		require.NoError(t, err)
		newBalance2, err := strconv.ParseFloat(result.ToAccount.Balance, 64)
		require.NoError(t, err)

		balanceDiffAccount1 := oldBalance1 - newBalance1
		balanceDiffAccount2 := newBalance2 - oldBalance2

		require.Equal(t, balanceDiffAccount1, balanceDiffAccount2)
		require.True(t, balanceDiffAccount1 > 0)
	}

	// get updated accounts
	updatedAccount1, err := store.GetAccount(context.Background(), createdAccount1.ID)
	require.NoError(t, err)
	updatedAccount2, err := store.GetAccount(context.Background(), createdAccount2.ID)
	require.NoError(t, err)

	amountF, err := strconv.ParseFloat(amount, 64)
	require.NoError(t, err)

	newBalance1, err := strconv.ParseFloat(updatedAccount1.Balance, 64)
	require.NoError(t, err)
	newBalance2, err := strconv.ParseFloat(updatedAccount2.Balance, 64)
	require.NoError(t, err)

	// verify final balances
	require.Equal(t, oldBalance1-float64(n)*amountF, newBalance1)
	require.Equal(t, oldBalance2+float64(n)*amountF, newBalance2)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := repository.NewStore(testDB)

	createdAccount1 := newAccount(t)
	createdAccount2 := newAccount(t)

	// number of transfers
	n := 10
	// amount to be transferred
	amount := "10.50"

	errs := make(chan error)

	// goroutines for transfers
	for i := 0; i < n; i++ {
		fromAccountID := createdAccount1.ID
		toAccountID := createdAccount2.ID

		if i%2 == 0 {
			fromAccountID = createdAccount2.ID
			toAccountID = createdAccount1.ID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), repository.TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	// checking results from goroutines
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

	}

	// get updated accounts
	updatedAccount1, err := store.GetAccount(context.Background(), createdAccount1.ID)
	require.NoError(t, err)
	updatedAccount2, err := store.GetAccount(context.Background(), createdAccount2.ID)
	require.NoError(t, err)

	// verify final balances
	require.Equal(t, createdAccount1.Balance, updatedAccount1.Balance)
	require.Equal(t, createdAccount2.Balance, updatedAccount2.Balance)
}

func newAccount(t *testing.T) sqlc.Account {
	userArg := devutils.RandomCreateUser()

	user, err := testQueries.CreateUser(context.Background(), userArg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	accountArg := devutils.RandomCreateAccount(user.Username)
	account, err := testQueries.CreateAccount(context.Background(), accountArg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}
