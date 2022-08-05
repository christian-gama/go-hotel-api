// config testando xd.
package config

import (
	"strconv"
	"time"

	"github.com/christian-gama/go-booking-api/internal/shared/infra/configger"
)

const (
	dbName               = "DB_NAME"
	dbUser               = "DB_USER"
	dbPassword           = "DB_PASSWORD"
	dbHost               = "DB_HOST"
	dbPort               = "DB_PORT"
	dbSslMode            = "DB_SSL_MODE"
	dbSgbd               = "DB_SGBD"
	dbMaxConnections     = "DB_MAX_CONNECTIONS"
	dbMaxIdleConnections = "DB_MAX_IDLE_CONNECTIONS"
	dbMaxLifeTimeMin     = "DB_MAX_LIFETIME_MIN"
	dbTimeoutSec         = "DB_TIMEOUT_SEC"
)

type db struct{}

// Name returns the name of the database.
func (d *db) Name() string {
	return getEnv(dbName)
}

// User returns the username of the database.
func (d *db) User() string {
	return getEnv(dbUser)
}

// Password returns the password of the database.
func (d *db) Password() string {
	return getEnv(dbPassword)
}

// Host returns the hostname of the database.
func (d *db) Host() string {
	return getEnv(dbHost)
}

// Port returns the port of the database.
func (d *db) Port() int {
	i, err := strconv.Atoi(getEnv(dbPort))
	if err != nil {
		panic(err)
	}
	return i
}

// SslMode returns the ssl mode of the database.
func (d *db) SslMode() string {
	return getEnv(dbSslMode)
}

// Sgbd returns which sgbd is being used.
func (d *db) Sgbd() string {
	return getEnv(dbSgbd)
}

// MaxConnections returns the maximum number of connections to the database.
func (d *db) MaxConnections() int {
	i, err := strconv.Atoi(getEnv(dbMaxConnections))
	if err != nil {
		panic(err)
	}
	return i
}

// MaxIdleConnections returns the maximum number of idle connections to the database.
func (d *db) MaxIdleConnections() int {
	i, err := strconv.Atoi(getEnv(dbMaxIdleConnections))
	if err != nil {
		panic(err)
	}
	return i
}

// MaxLifeTimeMin returns the maximum lifetime of a connection in minutes to the database.
func (d *db) MaxLifeTimeMin() time.Duration {
	i, err := strconv.Atoi(getEnv(dbMaxLifeTimeMin))
	if err != nil {
		panic(err)
	}
	return time.Duration(i) * time.Minute
}

// TimeoutSec returns the timeout in seconds of the database.
func (d *db) TimeoutSec() time.Duration {
	i, err := strconv.Atoi(getEnv(dbTimeoutSec))
	if err != nil {
		panic(err)
	}
	return time.Duration(i) * time.Second
}

// NewDb returns the database configuration.
func NewDb() configger.Db {
	return &db{}
}
