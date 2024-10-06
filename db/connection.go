package db

import (
	"database/sql"

	"github.com/Luks17/Go-Microservices-MC/db/repository"
)

var DBStore repository.Store

func InitStore(db *sql.DB) {
	DBStore = repository.NewStore(db)
}
