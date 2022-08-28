package factory

import (
	"database/sql"

	"github.com/christian-gama/go-hotel-api/internal/shared/infra/config"
	"github.com/christian-gama/go-hotel-api/internal/shared/ui/conn"
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
