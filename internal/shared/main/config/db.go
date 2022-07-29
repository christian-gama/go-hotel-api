package config

import (
	"strconv"
	"time"

	"github.com/christian-gama/go-booking-api/internal/shared/infra/configger"
)

type Db struct{}

func (d *Db) Name() string {
	return getEnv("DB_NAME")
}

func (d *Db) User() string {
	return getEnv("DB_USER")
}

func (d *Db) Password() string {
	return getEnv("DB_PASSWORD")
}

func (d *Db) Host() string {
	return getEnv("DB_HOST")
}

func (d *Db) Port() int {
	i, err := strconv.Atoi(getEnv("DB_PORT"))
	if err != nil {
		panic(err)
	}
	return i
}

func (d *Db) SslMode() string {
	return getEnv("DB_SSL_MODE")
}

func (d *Db) Sgbd() string {
	return getEnv("DB_SGDB")
}

func (d *Db) MaxConnections() int {
	i, err := strconv.Atoi(getEnv("DB_MAX_CONNECTIONS"))
	if err != nil {
		panic(err)
	}
	return i
}

func (d *Db) MaxIdleConnections() int {
	i, err := strconv.Atoi(getEnv("DB_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		panic(err)
	}
	return i
}

func (d *Db) MaxLifeTimeMin() time.Duration {
	i, err := strconv.Atoi(getEnv("DB_MAX_LIFETIME_MIN"))
	if err != nil {
		panic(err)
	}
	return time.Duration(i) * time.Minute
}

func (d *Db) TimeoutSec() time.Duration {
	i, err := strconv.Atoi(getEnv("DB_GET_TIMEOUT_SEC"))
	if err != nil {
		panic(err)
	}
	return time.Duration(i) * time.Second
}

func NewDb() configger.Db {
	return &Db{}
}
