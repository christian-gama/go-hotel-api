package factory

import (
	"database/sql"

	"github.com/christian-gama/go-booking-api/internal/infra/config"
	"github.com/christian-gama/go-booking-api/internal/ui/conn"
)

// PsqlConfig is a factory function that returns a new Postgres connection.
func PsqlConn() *sql.DB {
	dbConfig := config.NewDb()
	db, err := conn.NewSQL("pgx", dbConfig)
	if err != nil {
		panic(err)
	}

	return db
}
