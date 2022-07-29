package configger

import "time"

type App interface {
	Env() string
	Host() string
	Port() int
}

type Db interface {
	Name() string
	User() string
	Password() string
	Host() string
	Port() int
	SslMode() string
	Sgbd() string
	MaxConnections() int
	MaxIdleConnections() int
	MaxLifeTimeMin() time.Duration
	TimeoutSec() time.Duration
}
