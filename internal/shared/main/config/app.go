package config

import (
	"strconv"

	"github.com/christian-gama/go-booking-api/internal/shared/infra/configger"
)

type app struct{}

func (a *app) Env() string {
	return getEnv("ENV")
}

func (a *app) Host() string {
	return getEnv("APP_HOST")
}

func (a *app) Port() int {
	i, err := strconv.Atoi(getEnv("APP_PORT"))
	if err != nil {
		panic(err)
	}
	return i
}

func NewApp() configger.App {
	return &app{}
}
