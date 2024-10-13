package sqlc_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Luks17/Go-Microservices-MC/devutils"
	"github.com/Luks17/Go-Microservices-MC/devutils/devmodels"
	"github.com/stretchr/testify/require"
)

func TestCreateGetUser(t *testing.T) {
	password := devutils.RandomPassword(t)

	arg := devmodels.RandomUserParams(password)
	createdUser, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, createdUser)

	require.Equal(t, arg.Username, createdUser.Username)
	require.Equal(t, arg.Password, createdUser.Password)
	require.Equal(t, arg.FullName, createdUser.FullName)
	require.Equal(t, arg.Email, createdUser.Email)
	require.NotZero(t, createdUser.CreatedAt)
	require.IsType(t, sql.NullTime{}, createdUser.PasswordLastChangedAt)

	retrievedUser, err := testQueries.GetUser(context.Background(), createdUser.Username)

	require.NoError(t, err)
	require.NotEmpty(t, retrievedUser)

	require.Equal(t, createdUser.Username, retrievedUser.Username)
	require.Equal(t, createdUser.Password, retrievedUser.Password)
	require.Equal(t, createdUser.FullName, retrievedUser.FullName)
	require.Equal(t, createdUser.Email, retrievedUser.Email)
	require.WithinDuration(t, createdUser.CreatedAt, retrievedUser.CreatedAt, time.Second)
	require.IsType(t, sql.NullTime{}, retrievedUser.PasswordLastChangedAt)
}
