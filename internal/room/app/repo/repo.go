package repo

import (
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
)

// Room is the interface for the room repo.
type Room interface {
	// SaveRoom is the method that will save a room.
	SaveRoom(room *entity.Room) (*entity.Room, []*errorutil.Error)

	// GetRoom is the method that will get a room by its uuid.
	GetRoom(uuid string) (*entity.Room, []*errorutil.Error)

	// ListRooms is the method that will list all rooms.
	ListRooms() ([]*entity.Room, []*errorutil.Error)
}
