package factory

import (
	"github.com/christian-gama/go-booking-api/internal/room/infra/repo/psql"
	sharedfty "github.com/christian-gama/go-booking-api/internal/shared/ui/factory"
)

// RoomRepo is a factory function that returns a new room repository.
func RoomRepo() psql.RoomRepo {
	conn := sharedfty.PsqlConn()
	dbConfig := sharedfty.DbConfig()

	return psql.NewRoomRepo(conn, dbConfig)
}
