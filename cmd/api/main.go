package main

import (
	"fmt"
	"net/http"

	"github.com/christian-gama/go-booking-api/internal/main/factory"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	listenAndServe()
}

func listenAndServe() {
	appConfig := factory.AppConfig()
	r := factory.Router()

	fmt.Printf("Server listening on %d\n", appConfig.Port())
	err := http.ListenAndServe(fmt.Sprintf(":%d", appConfig.Port()), r.Mux)
	if err != nil {
		panic(fmt.Errorf("cannot start application: %w", err))
	}
}
