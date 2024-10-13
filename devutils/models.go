package devutils

import (
	"context"
	"testing"

	"github.com/Luks17/Go-Microservices-MC/db/sqlc"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

func RandomUserParams(password string) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Username: gofakeit.Username(),
		Password: password,
		FullName: gofakeit.Name(),
		Email:    gofakeit.Email(),
	}
}

func CreateNewRandomUser(t *testing.T, queries *sqlc.Queries) sqlc.User {
	password := RandomPassword(t)

	arg := RandomUserParams(password)
	user, err := queries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}

func RandomAccountParams(username string) sqlc.CreateAccountParams {
	return sqlc.CreateAccountParams{
		Owner:    username,
		Balance:  RandomBalance(),
		Currency: RandomCurrency(),
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
		Balance:   RandomBalance(),
		Currency:  RandomCurrency(),
		CreatedAt: RandomTimeStamp(),
	}
}
