package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:123456@0.0.0.0:5432/bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}

func RandomCurrency() Currencies {
	currencies := []Currencies{CurrenciesUSD, CurrenciesEUR, CurrenciesBRL}

	return currencies[gofakeit.Number(0, len(currencies)-1)]
}
