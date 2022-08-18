package config

import (
	"strconv"
)

const (
	env     = "ENV"
	appHost = "APP_HOST"
	appPort = "APP_PORT"
)

type (
	// App is the application configuration.
	App interface {
		Env() string
		Host() string
		Port() int
	}

	// appImpl is the implementation of the App interface.
	appImpl struct{}
)

// Env returns the environment name (dev, prod, test).
func (a *appImpl) Env() string {
	return getEnv(env)
}

// Host returns the hostname of the application.
func (a *appImpl) Host() string {
	return getEnv(appHost)
}

// Port returns the port of the application.
func (a *appImpl) Port() int {
	i, err := strconv.Atoi(getEnv(appPort))
	if err != nil {
		panic(err)
	}
	return i
}

// NewApp returns the application configuration.
func NewApp() App {
	return &appImpl{}
}
