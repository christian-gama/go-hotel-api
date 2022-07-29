package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/christian-gama/go-booking-api/internal/shared/infra/configger"
	"github.com/christian-gama/go-booking-api/internal/shared/main/config"
	"github.com/christian-gama/go-booking-api/internal/shared/main/conn"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	appConfig configger.App
	dbConfig  configger.Db
	db        *sql.DB
)

func main() {
	setupEnv()
	setupDb()
	listenAndServe()
}

func setupEnv() {
	config.LoadEnvFile(".env.dev")
	appConfig = config.NewApp()
	dbConfig = config.NewDb()

	log.Printf("Starting server on %s environment...\n", appConfig.Env())
}

func setupDb() {
	log.Printf("Connecting to database...\n")
	var err error
	db, err = conn.NewSQL("pgx", dbConfig)
	if err != nil {
		panic(err)
	}
}

func listenAndServe() {
	defer db.Close()

	log.Printf("Listening on: http://%s:%d\n", appConfig.Host(), appConfig.Port())

	// get route "/"
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Hello, world!")
		fmt.Fprintf(w, "Hello, world!")
	})

	err := http.ListenAndServe(fmt.Sprintf(":%d", appConfig.Port()), nil)
	if err != nil {
		panic(fmt.Errorf("cannot start application: %w", err))
	}
}
