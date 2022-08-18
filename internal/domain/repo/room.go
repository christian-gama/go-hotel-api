package repo

import (
	"github.com/christian-gama/go-booking-api/internal/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/domain/error"
)

type (
	// SaveRoomRepo is the interface for saving a room.
	SaveRoomRepo interface {
		SaveRoom(room *entity.Room) (*entity.Room, []*error.Error)
	}

	// DeleteRoomRepo is the interface for deleting a room.
	DeleteRoomRepo interface {
		DeleteRoom(uuid string) (bool, []*error.Error)
	}

	// GetRoomRepo is the interface for getting a room.
	GetRoomRepo interface {
		GetRoom(uuid string) (*entity.Room, []*error.Error)
	}

	// ListRoomsRepo is the interface for listing all rooms.
	ListRoomsRepo interface {
		ListRooms() ([]*entity.Room, []*error.Error)
	}
)
