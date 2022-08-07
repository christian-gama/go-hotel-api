package config

import (
	"strconv"

	"github.com/christian-gama/go-booking-api/internal/shared/infra/configger"
)

const (
	env     = "ENV"
	appHost = "APP_HOST"
	appPort = "APP_PORT"
)

type app struct{}

// Env returns the environment name (dev, prod, test).
func (a *app) Env() string {
	return getEnv(env)
}

// Host returns the hostname of the application.
func (a *app) Host() string {
	return getEnv(appHost)
}

// Port returns the port of the application.
func (a *app) Port() int {
	i, err := strconv.Atoi(getEnv(appPort))
	if err != nil {
		panic(err)
	}
	return i
}

// NewApp returns the application configuration.
func NewApp() configger.App {
	return &app{}
}
