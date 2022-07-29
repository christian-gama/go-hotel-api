package repo

import (
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/app/exception"
)

// Room is the interface for the room repo.
type Room interface {
	SaveRoom(room *entity.Room) (*entity.Room, *exception.Error)
}
