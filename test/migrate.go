package test

import (
	"database/sql"
	"fmt"

	"github.com/christian-gama/go-booking-api/internal/shared/main/config"
	"github.com/christian-gama/go-booking-api/internal/shared/util"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// sources the postgres driver.
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(db *sql.DB) *migrate.Migrate {
	migConfigger := config.NewMig()

	driver, err := postgres.WithInstance(db, &postgres.Config{DatabaseName: "gobooking_test"})
	if err != nil {
		panic(err)
	}

	fmt.Println(util.RootPath())
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s/%s", util.RootPath(), migConfigger.Dir()),
		"postgres", driver)
	if err != nil {
		panic(err)
	}

	return m
}
