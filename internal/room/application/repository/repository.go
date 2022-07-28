package repository

import "github.com/christian-gama/go-booking-api/internal/room/domain"

type Repository interface {
	SaveRoom(room *domain.Room) (*domain.Room, error)
}
