package internal

import (
	"fmt"
	"log"
	"net/http"

	roomRoutes "github.com/christian-gama/go-booking-api/internal/room/ui/routes"
	"github.com/christian-gama/go-booking-api/internal/shared/infra/config"
	"github.com/christian-gama/go-booking-api/internal/shared/infra/router"
	"github.com/christian-gama/go-booking-api/internal/shared/ui/factory"
	sharedRoutes "github.com/christian-gama/go-booking-api/internal/shared/ui/routes"
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
	router.RegisterAll(
		roomRoutes.New(),
		sharedRoutes.New(),
	)

	return router
}

func setupConfig() (config.App, config.Db) {
	appConfig := factory.AppConfig()
	dbConfig := factory.DbConfig()

	return appConfig, dbConfig
}
