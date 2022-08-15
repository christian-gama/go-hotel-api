package conn

import (
	_sql "database/sql"
	"fmt"
	"time"

	"github.com/christian-gama/go-booking-api/internal/infra/config"
)

type sql struct {
	dbConfig config.Db

	db         *_sql.DB
	driverName string
	dsn        string
}

// open opens a connection to the database. It will return an error if fails to connect.
func (s *sql) open() error {
	s.setDsn()

	db, err := _sql.Open(s.driverName, s.dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	const attempts = 3
	for i := 0; i <= attempts; i++ {
		err = db.Ping()
		if err == nil {
			break
		}

		if i == attempts {
			return fmt.Errorf("failed to ping database after %d attempts: %w", i, err)
		}

		time.Sleep(time.Second * 2)
	}

	s.db = db

	return nil
}

// setup sets up the database configuration, such as max connections, idle connections, etc.
func (s *sql) setup() {
	s.db.SetMaxIdleConns(s.dbConfig.MaxIdleConnections())
	s.db.SetMaxOpenConns(s.dbConfig.MaxConnections())
	s.db.SetConnMaxLifetime(s.dbConfig.MaxLifeTime())
}

// setDsn sets the data source name.
func (s *sql) setDsn() {
	s.dsn = fmt.Sprintf(`
	host=%s
	port=%d
	dbname=%s
	user=%s
	password=%s
	sslmode=%s`,
		s.dbConfig.Host(),
		s.dbConfig.Port(),
		s.dbConfig.Name(),
		s.dbConfig.User(),
		s.dbConfig.Password(),
		s.dbConfig.SslMode(),
	)
}

// NewSQL returns a new sql connection.
func NewSQL(driverName string, dbConfig config.Db) (*_sql.DB, error) {
	sql := &sql{
		dbConfig:   dbConfig,
		driverName: driverName,
	}

	if err := sql.open(); err != nil {
		return nil, err
	}

	sql.setup()

	return sql.db, nil
}
