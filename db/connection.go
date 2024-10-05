package db

import (
	"database/sql"

	"github.com/Luks17/Go-Microservices-MC/db/tx"
)

var DBStore *tx.Store

func InitStore(db *sql.DB) {
	DBStore = tx.NewStore(db)
}
