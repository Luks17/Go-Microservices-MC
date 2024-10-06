package db

import (
	"database/sql"

	"github.com/Luks17/Go-Microservices-MC/db/repository"
)

var DBStore *repository.SQLStore

func InitStore(db *sql.DB) {
	DBStore = repository.NewStore(db)
}
