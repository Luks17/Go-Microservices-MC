package devutils

import (
	"github.com/Luks17/Go-Microservices-MC/db/sqlc"
	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/lib/pq"
)

func RandomCurrency() sqlc.Currencies {
	currencies := []sqlc.Currencies{sqlc.CurrenciesUSD, sqlc.CurrenciesEUR, sqlc.CurrenciesBRL}

	return currencies[gofakeit.Number(0, len(currencies)-1)]
}
