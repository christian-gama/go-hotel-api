package service

import (
	"github.com/christian-gama/go-booking-api/internal/room/application/repository"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
)

type CreateRoomService interface {
	Handle(input *entity.Room) (*entity.Room, error)
}

type CreateRoom struct {
	repository repository.Repository
}

func (c *CreateRoom) Handle(input *entity.Room) (*entity.Room, error) {
	room, err := c.repository.SaveRoom(input)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func NewCreateRoom(repository repository.Repository) CreateRoomService {
	return &CreateRoom{
		repository: repository,
	}
}
