package usecase

import (
	"github.com/christian-gama/go-booking-api/internal/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/domain/repo"
)

type (
	// GetRoomUsecase is the interface that defines the retrieval of a room.
	GetRoomUsecase interface {
		Handle(uuid string) (*entity.Room, []*errorutil.Error)
	}

	// getRoomImpl is a concrete implementation of the GetRoom.
	getRoomImpl struct {
		repo repo.GetRoomRepo
	}
)

// Handle receives a uuid and retrieves a room. It will return an error if something
// goes wrong with room retrieval or if the room repo return an error.
func (g *getRoomImpl) Handle(uuid string) (*entity.Room, []*errorutil.Error) {
	room, err := g.repo.GetRoom(uuid)
	if err != nil {
		return nil, err
	}

	return room, nil
}

// NewGetRoom creates a new GetRoom.
func NewGetRoom(repo repo.GetRoomRepo) GetRoomUsecase {
	return &getRoomImpl{
		repo,
	}
}
