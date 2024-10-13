package devutils

import (
	"strconv"
	"testing"
	"time"

	"github.com/Luks17/Go-Microservices-MC/crypt"
	"github.com/Luks17/Go-Microservices-MC/db/sqlc"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func RandomCurrency() sqlc.Currencies {
	currencies := []sqlc.Currencies{sqlc.CurrenciesUSD, sqlc.CurrenciesEUR, sqlc.CurrenciesBRL}

	return currencies[gofakeit.Number(0, len(currencies)-1)]
}

func RandomTimeStamp() time.Time {
	return gofakeit.DateRange(time.Now().AddDate(-5, 0, 0), time.Now())
}

func RandomBalance() string {
	return strconv.FormatFloat(gofakeit.Price(0, 10000), 'f', 2, 64)
}

func RandomPassword(t *testing.T) string {
	password := gofakeit.Password(true, true, true, false, false, 12)
	hashedPassword, err := crypt.HashPassword(password)

	require.NoError(t, err)

	return hashedPassword
}
