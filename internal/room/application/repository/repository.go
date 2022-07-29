package repository

import (
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/application/exception"
)

// Room is the interface for the room repository.
type Room interface {
	SaveRoom(room *entity.Room) (*entity.Room, *exception.Error)
}
