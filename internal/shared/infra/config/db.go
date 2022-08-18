// config testando xd.
package config

import (
	"strconv"
	"time"
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

type (
	// Db is the database configuration.
	Db interface {
		Name() string
		User() string
		Password() string
		Host() string
		Port() int
		SslMode() string
		Sgbd() string
		MaxConnections() int
		MaxIdleConnections() int
		MaxLifeTime() time.Duration
		Timeout() time.Duration
	}

	// dbImpl is the implementation of the Db interface.
	dbImpl struct{}
)

// Name returns the name of the database.
func (d *dbImpl) Name() string {
	return getEnv(dbName)
}

// User returns the username of the database.
func (d *dbImpl) User() string {
	return getEnv(dbUser)
}

// Password returns the password of the database.
func (d *dbImpl) Password() string {
	return getEnv(dbPassword)
}

// Host returns the hostname of the database.
func (d *dbImpl) Host() string {
	return getEnv(dbHost)
}

// Port returns the port of the database.
func (d *dbImpl) Port() int {
	i, err := strconv.Atoi(getEnv(dbPort))
	if err != nil {
		panic(err)
	}
	return i
}

// SslMode returns the ssl mode of the database.
func (d *dbImpl) SslMode() string {
	return getEnv(dbSslMode)
}

// Sgbd returns which sgbd is being used.
func (d *dbImpl) Sgbd() string {
	return getEnv(dbSgbd)
}

// MaxConnections returns the maximum number of connections to the database.
func (d *dbImpl) MaxConnections() int {
	i, err := strconv.Atoi(getEnv(dbMaxConnections))
	if err != nil {
		panic(err)
	}
	return i
}

// MaxIdleConnections returns the maximum number of idle connections to the database.
func (d *dbImpl) MaxIdleConnections() int {
	i, err := strconv.Atoi(getEnv(dbMaxIdleConnections))
	if err != nil {
		panic(err)
	}
	return i
}

// MaxLifeTime returns the maximum lifetime of a connection in minutes to the database.
func (d *dbImpl) MaxLifeTime() time.Duration {
	i, err := strconv.Atoi(getEnv(dbMaxLifeTimeMin))
	if err != nil {
		panic(err)
	}
	return time.Duration(i) * time.Minute
}

// Timeout returns the timeout in seconds of the database.
func (d *dbImpl) Timeout() time.Duration {
	i, err := strconv.Atoi(getEnv(dbTimeoutSec))
	if err != nil {
		panic(err)
	}
	return time.Duration(i) * time.Second
}

// NewDb returns the database configuration.
func NewDb() Db {
	return &dbImpl{}
}
