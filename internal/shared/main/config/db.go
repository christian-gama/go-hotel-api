package config

import (
	"strconv"
	"time"

	"github.com/christian-gama/go-booking-api/internal/shared/infra/configger"
)

type db struct{}

// Name returns the name of the database.
func (d *db) Name() string {
	return getEnv("DB_NAME")
}

// User returns the username of the database.
func (d *db) User() string {
	return getEnv("DB_USER")
}

// Password returns the password of the database.
func (d *db) Password() string {
	return getEnv("DB_PASSWORD")
}

// Host returns the hostname of the database.
func (d *db) Host() string {
	return getEnv("DB_HOST")
}

// Port returns the port of the database.
func (d *db) Port() int {
	i, err := strconv.Atoi(getEnv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	return i
}

// SslMode returns the ssl mode of the database.
func (d *db) SslMode() string {
	return getEnv("DB_SSL_MODE")
}

// Sgbd returns which sgbd is being used.
func (d *db) Sgbd() string {
	return getEnv("DB_SGDB")
}

// MaxConnections returns the maximum number of connections to the database.
func (d *db) MaxConnections() int {
	i, err := strconv.Atoi(getEnv("DB_MAX_CONNECTIONS"))
	if err != nil {
		panic(err)
	}
	return i
}

// MaxIdleConnections returns the maximum number of idle connections to the database.
func (d *db) MaxIdleConnections() int {
	i, err := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		panic(err)
	}
	return i
}

// MaxLifeTimeMin returns the maximum lifetime of a connection in minutes to the database.
func (d *db) MaxLifeTimeMin() time.Duration {
	i, err := strconv.Atoi(getEnv("DB_MAX_LIFETIME_MIN"))
	if err != nil {
		panic(err)
	}
	return time.Duration(i) * time.Minute
}

// TimeoutSec returns the timeout in seconds of the database.
func (d *db) TimeoutSec() time.Duration {
	i, err := strconv.Atoi(getEnv("DB_GET_TIMEOUT_SEC"))
	if err != nil {
		panic(err)
	}
	return time.Duration(i) * time.Second
}

// NewDb returns the database configuration.
func NewDb() configger.Db {
	return &db{}
}
