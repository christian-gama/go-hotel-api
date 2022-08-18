package factory

import (
	"github.com/christian-gama/go-booking-api/internal/infra/repo/psql"
)

// RoomRepo is a factory function that returns a new room repository.
func RoomRepo() psql.RoomRepo {
	conn := PsqlConn()
	dbConfig := DbConfig()

	return psql.NewRoomRepo(conn, dbConfig)
}
