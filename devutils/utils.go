package devutils

import (
	"strconv"
	"time"

	"github.com/Luks17/Go-Microservices-MC/db/sqlc"
	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/lib/pq"
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

func RandomCreateAccount() sqlc.CreateAccountParams {
	return sqlc.CreateAccountParams{
		Owner:    gofakeit.Name(),
		Balance:  RandomBalance(),
		Currency: RandomCurrency(),
	}
}

func RandomNewAccount() sqlc.Account {
	return sqlc.Account{
		ID:        gofakeit.Int64(),
		Owner:     gofakeit.Name(),
		Balance:   RandomBalance(),
		Currency:  RandomCurrency(),
		CreatedAt: RandomTimeStamp(),
	}
}
