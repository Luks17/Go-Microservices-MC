package crypt_test

import (
	"testing"

	"github.com/Luks17/Go-Microservices-MC/crypt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestPassword(t *testing.T) {
	password := gofakeit.Password(true, true, true, false, false, 16)

	hashedPassword1, err := crypt.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = crypt.CheckPassword(password, hashedPassword1)
	require.NoError(t, err)

	hashedPassword2, err := crypt.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	// even though the passwords are the same, the hashes must be different
	require.NotEqual(t, hashedPassword1, hashedPassword2)

	wrongPassword := gofakeit.Password(true, true, true, false, false, 16)
	err = crypt.CheckPassword(wrongPassword, hashedPassword1)
	require.EqualError(t, err, crypt.ErrPasswordsDoNotMatch.Error())
}
