package ui

import (
	"fmt"
	"log"
	"net/http"

	"github.com/christian-gama/go-booking-api/internal/infra/config"
	"github.com/christian-gama/go-booking-api/internal/infra/router"
	"github.com/christian-gama/go-booking-api/internal/ui/factory"
	"github.com/christian-gama/go-booking-api/internal/ui/routes"
)

// Bootstrap is the main function that starts the application.
func Bootstrap() {
	appConfig, _ := setupConfig()
	router := setupRouter()

	log.Printf("Server is running on port :%d", appConfig.Port())
	err := http.ListenAndServe(fmt.Sprintf(":%d", appConfig.Port()), router.Mux)
	if err != nil {
		log.Panicf("Error starting server: %s", err.Error())
	}
}

func setupRouter() *router.Router {
	router := factory.Router()
	defer router.Load()

	routes.Register(router)

	return router
}

func setupConfig() (config.App, config.Db) {
	appConfig := factory.AppConfig()
	dbConfig := factory.DbConfig()

	return appConfig, dbConfig
}
