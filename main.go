package main

import (
	"database/sql"
	"log"

	"github.com/Luks17/Go-Microservices-MC/api"
	"github.com/Luks17/Go-Microservices-MC/db"
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:123456@0.0.0.0:5432/bank?sslmode=disable"
	serverAddress = "0.0.0.0:7717"
)

func main() {
	dbConn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	db.InitStore(dbConn)

	err = api.InitServer(serverAddress)
	if err != nil {
		log.Fatal("Could not bind server to address: ", err)
	}
}
