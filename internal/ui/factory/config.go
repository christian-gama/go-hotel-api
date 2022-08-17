package factory

import (
	"github.com/christian-gama/go-booking-api/internal/infra/config"
)

func init() {
	config.LoadEnvFile(".env.dev")
}

// DbConfig is a factory function that returns a new database config.
func DbConfig() config.Db {
	return config.NewDb()
}

// AppConfig is a factory function that returns a new application config.
func AppConfig() config.App {
	return config.NewApp()
}
