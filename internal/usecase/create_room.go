package usecase

import (
	"github.com/christian-gama/go-booking-api/internal/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/domain/repo"
	"github.com/christian-gama/go-booking-api/internal/domain/uuid"
)

type (
	// CreateRoomInput reprensents the input of the CreateRoom.
	CreateRoomInput struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		BedCount    uint8   `json:"bed_count"`
		Price       float32 `json:"price"`
	}

	// CreateRoom is the interface that defines the creation of a room.
	CreateRoom interface {
		Handle(input *CreateRoomInput) (*entity.Room, []*errorutil.Error)
	}

	// createRoomImpl is a concrete implementation of the CreateRoom.
	createRoomImpl struct {
		repo repo.SaveRoom
		uuid uuid.UUID
	}
)

// Handle receives an input and creates a room. It will return an error if something
// goes wrong with room creation or if the room repo return an error.
func (c *createRoomImpl) Handle(input *CreateRoomInput) (*entity.Room, []*errorutil.Error) {
	uuid := c.uuid.Generate()
	room, err := entity.NewRoom(uuid, input.Name, input.Description, input.BedCount, input.Price)
	if err != nil {
		return nil, err
	}

	room, err = c.repo.SaveRoom(room)
	if err != nil {
		return nil, err
	}

	return room, nil
}

// NewCreateRoom creates a new CreateRoom.
func NewCreateRoom(repo repo.SaveRoom, uuid uuid.UUID) CreateRoom {
	return &createRoomImpl{
		repo,
		uuid,
	}
}
