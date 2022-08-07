package test

import (
	"database/sql"

	"github.com/christian-gama/go-booking-api/internal/shared/main/config"
	"github.com/christian-gama/go-booking-api/internal/shared/main/conn"

	// sources postgres driver.
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func SetupDb() *sql.DB {
	config.LoadEnvFile(".env.test")

	db, err := conn.NewSQL("pgx", config.NewDb())
	if err != nil {
		panic(err)
	}

	return db
}
