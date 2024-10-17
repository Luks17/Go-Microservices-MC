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

func RandomUserParams(password string) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Username: gofakeit.Username(),
		Password: password,
		FullName: gofakeit.Name(),
		Email:    gofakeit.Email(),
	}
}

func CreateNewRandomUser(t *testing.T, queries *sqlc.Queries) sqlc.User {
	password := devutils.RandomPassword(t)

	arg := RandomUserParams(password)
	user, err := queries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}
