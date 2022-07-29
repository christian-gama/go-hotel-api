package configger

import "time"

// App is the application configuration.
type App interface {
	Env() string
	Host() string
	Port() int
}

// Db is the database configuration.
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
