package factory

import (
	"database/sql"

	"github.com/christian-gama/go-booking-api/internal/infra/config"
	"github.com/christian-gama/go-booking-api/internal/main/conn"
)

func PsqlConn() *sql.DB {
	dbConfig := config.NewDb()
	db, err := conn.NewSQL("pgx", dbConfig)
	if err != nil {
		panic(err)
	}

	return db
}
