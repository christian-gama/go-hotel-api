package service

import (
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

// Handle receives an input and creates a room. It will return an error if something
// goes wrong with room creation or if the room repository return an error.
func (c *CreateRoom) Handle(input *CreateRoomInput) (*entity.Room, *exception.Error) {
	uuid := c.uuid.Generate()
	room, err := entity.NewRoom(uuid, input.Name, input.Description, input.BedCount, input.Price, false)
	if err != nil {
		return nil, c.exception.BadRequest(err.Error())
	}

	room, exception := c.repository.SaveRoom(room)
	if exception != nil {
		return nil, exception
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
