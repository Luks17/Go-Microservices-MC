package repository_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Luks17/Go-Microservices-MC/db/repository/sqlc"
	"github.com/Luks17/Go-Microservices-MC/util"
)

var (
	testQueries *sqlc.Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Could not load config: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, util.GetDBConnectionURI(&config))
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = sqlc.New(testDB)

	os.Exit(m.Run())
}
