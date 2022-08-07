package conn

import (
	_sql "database/sql"
	"fmt"
	"time"

	"github.com/christian-gama/go-booking-api/internal/shared/infra/configger"
)

type sql struct {
	dbConfigger configger.Db

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

	for i := 0; i <= 3; i++ {
		err = db.Ping()
		if err == nil {
			break
		}

		if i == 3 {
			return fmt.Errorf("failed to ping database: %w", err)
		}

		time.Sleep(time.Second * 2)
	}

	s.db = db

	return nil
}

// setup sets up the database configuration, such as max connections, idle connections, etc.
func (s *sql) setup() {
	s.db.SetMaxIdleConns(s.dbConfigger.MaxIdleConnections())
	s.db.SetMaxOpenConns(s.dbConfigger.MaxConnections())
	s.db.SetConnMaxLifetime(s.dbConfigger.MaxLifeTime())
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
		s.dbConfigger.Host(),
		s.dbConfigger.Port(),
		s.dbConfigger.Name(),
		s.dbConfigger.User(),
		s.dbConfigger.Password(),
		s.dbConfigger.SslMode(),
	)
}

// NewSQL returns a new sql connection.
func NewSQL(driverName string, dbConfigger configger.Db) (*_sql.DB, error) {
	sql := &sql{
		dbConfigger: dbConfigger,
		driverName:  driverName,
	}

	if err := sql.open(); err != nil {
		return nil, err
	}

	sql.setup()

	return sql.db, nil
}
