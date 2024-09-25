package db

import (
	"context"
	"database/sql"
	"strconv"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func randomNewAccount() CreateAccountParams {
	return CreateAccountParams{
		Owner:    gofakeit.Name(),
		Balance:  strconv.FormatFloat(gofakeit.Price(0, 10000), 'f', 2, 64),
		Currency: RandomCurrency(),
	}
}

func TestCreateAccount(t *testing.T) {
	arg := randomNewAccount()

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	arg := randomNewAccount()

	createdAccount, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, createdAccount)

	retrievedAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, retrievedAccount)

	require.Equal(t, createdAccount.ID, retrievedAccount.ID)
	require.Equal(t, createdAccount.Owner, retrievedAccount.Owner)
	require.Equal(t, createdAccount.Balance, retrievedAccount.Balance)
	require.Equal(t, createdAccount.Currency, retrievedAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, retrievedAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	createArg := randomNewAccount()

	createdAccount, err := testQueries.CreateAccount(context.Background(), createArg)

	require.NoError(t, err)
	require.NotEmpty(t, createdAccount)

	updateArg := UpdateAccountParams{
		ID:      createdAccount.ID,
		Balance: strconv.FormatFloat(gofakeit.Price(0, 10000), 'f', 2, 64),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), updateArg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, createdAccount.ID, updatedAccount.ID)
	require.Equal(t, createdAccount.Owner, updatedAccount.Owner)
	require.Equal(t, updateArg.Balance, updatedAccount.Balance)
	require.Equal(t, createdAccount.Currency, updatedAccount.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, updatedAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	createArg := randomNewAccount()

	createdAccount, err := testQueries.CreateAccount(context.Background(), createArg)

	require.NoError(t, err)
	require.NotEmpty(t, createdAccount)

	err = testQueries.DeleteAccount(context.Background(), createdAccount.ID)

	require.NoError(t, err)

	deletedAccount, err := testQueries.GetAccount(context.Background(), createdAccount.ID)

	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, deletedAccount)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createArg := randomNewAccount()

		createdAccount, err := testQueries.CreateAccount(context.Background(), createArg)

		require.NoError(t, err)
		require.NotEmpty(t, createdAccount)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
