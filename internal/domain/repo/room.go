package repo

import (
	"github.com/christian-gama/go-booking-api/internal/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
)

type (
	// SaveRoom is the interface for saving a room.
	SaveRoom interface {
		SaveRoom(room *entity.Room) (*entity.Room, []*errorutil.Error)
	}

	// DeleteRoom is the interface for deleting a room.
	DeleteRoom interface {
		DeleteRoom(uuid string) []*errorutil.Error
	}

	// GetRoom is the interface for getting a room.
	GetRoom interface {
		GetRoom(uuid string) (*entity.Room, []*errorutil.Error)
	}

	// ListRooms is the interface for listing all rooms.
	ListRooms interface {
		ListRooms() ([]*entity.Room, []*errorutil.Error)
	}
)
