package repository_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Luks17/Go-Microservices-MC/db/sqlc"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:123456@0.0.0.0:5432/bank?sslmode=disable"
)

var (
	testQueries *sqlc.Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = sqlc.New(testDB)

	os.Exit(m.Run())
}
