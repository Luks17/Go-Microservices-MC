package sqlc_test

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Luks17/Go-Microservices-MC/db/sqlc"
	"github.com/Luks17/Go-Microservices-MC/util"
)

var testQueries *sqlc.Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Could not load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, util.GetDBConnectionURI(&config))
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = sqlc.New(conn)

	os.Exit(m.Run())
}
