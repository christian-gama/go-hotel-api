package service

import (
	"fmt"

	"github.com/christian-gama/go-booking-api/internal/room/application/repository"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/application/exception"
	"github.com/christian-gama/go-booking-api/internal/shared/application/uuid"
)

type CreateRoom struct {
	repository repository.Room
	exception  exception.Exception
	uuid       uuid.UUID
}

type CreateRoomInput struct {
	Name        string
	Description string
	BedCount    uint8
	Price       float32
}

func (c *CreateRoom) Handle(input *CreateRoomInput) (*entity.Room, *exception.Error) {
	uuid := c.uuid.Generate()
	fmt.Println(uuid)
	room, err := entity.NewRoom(uuid, input.Name, input.Description, input.BedCount, input.Price, false)
	if err != nil {
		return nil, c.exception.BadRequest(err.Error())
	}

	room, err = c.repository.SaveRoom(room)
	if err != nil {
		return nil, c.exception.BadRequest(err.Error())
	}

	return room, nil
}

func NewCreateRoom(repository repository.Room, exception exception.Exception, uuid uuid.UUID) *CreateRoom {
	return &CreateRoom{
		repository,
		exception,
		uuid,
	}
}
