package sqlc_test

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

var testQueries *sqlc.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = sqlc.New(conn)

	os.Exit(m.Run())
}
