package repository

import "github.com/christian-gama/go-booking-api/internal/room/domain/entity"

type Repository interface {
	SaveRoom(room *entity.Room) (*entity.Room, error)
}
