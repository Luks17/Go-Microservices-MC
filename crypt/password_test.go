package crypt_test

import (
	"testing"

	"github.com/Luks17/Go-Microservices-MC/crypt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := gofakeit.Password(true, true, true, false, false, 16)

	hashedPassword, err := crypt.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = crypt.CheckPassword(password, hashedPassword)
	require.NoError(t, err)

	wrongPassword := gofakeit.Password(true, true, true, false, false, 16)
	err = crypt.CheckPassword(wrongPassword, hashedPassword)
	require.EqualError(t, err, crypt.ErrPasswordsDoNotMatch.Error())
}
