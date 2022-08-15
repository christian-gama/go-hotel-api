package factory

import (
	"github.com/christian-gama/go-booking-api/internal/infra/config"
)

func init() {
	config.LoadEnvFile(".env.dev")
}

func DbConfig() config.Db {
	return config.NewDb()
}

func AppConfig() config.App {
	return config.NewApp()
}
