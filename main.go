package main

import (
	"database/sql"
	"log"

	"github.com/Luks17/Go-Microservices-MC/api"
	"github.com/Luks17/Go-Microservices-MC/db"
	"github.com/Luks17/Go-Microservices-MC/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config: ", err)
	}

	dbConn, err := sql.Open(config.DBDriver, util.GetDBConnectionURI(&config))
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	db.InitStore(dbConn)

	err = api.InitServer(util.GetServerURI(&config))
	if err != nil {
		log.Fatal("Could not bind server to address: ", err)
	}
}
