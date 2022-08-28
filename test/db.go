package test

import (
	"database/sql"

	"github.com/christian-gama/go-hotel-api/internal/shared/infra/config"
	"github.com/christian-gama/go-hotel-api/internal/shared/ui/conn"
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
