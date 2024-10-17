package devmodels

import (
	"context"
	"testing"

	"github.com/Luks17/Go-Microservices-MC/db/repository/sqlc"
	"github.com/Luks17/Go-Microservices-MC/devutils"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

func RandomAccountParams(username string) sqlc.CreateAccountParams {
	return sqlc.CreateAccountParams{
		Owner:    username,
		Balance:  devutils.RandomBalance(),
		Currency: devutils.RandomCurrency(),
	}
}

func CreateNewRandomAccount(t *testing.T, queries *sqlc.Queries) sqlc.Account {
	user := CreateNewRandomUser(t, queries)

	arg := RandomAccountParams(user.Username)
	account, err := queries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	return account
}

func RandomMockAccount() sqlc.Account {
	return sqlc.Account{
		ID:        gofakeit.Int64(),
		Owner:     gofakeit.Username(),
		Balance:   devutils.RandomBalance(),
		Currency:  devutils.RandomCurrency(),
		CreatedAt: devutils.RandomTimeStamp(),
	}
}
