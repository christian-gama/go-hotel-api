package usecase

import (
	"github.com/christian-gama/go-booking-api/internal/room/app/protocol"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
)

// GetRoom is the interface that defines the retrieval of a room.
type GetRoom interface {
	Handle(uuid string) (*entity.Room, []*errorutil.Error)
}

// getRoom is a concrete implementation of the GetRoom.
type getRoom struct {
	repo protocol.GetRoomRepo
}

// Handle receives a uuid and retrieves a room. It will return an error if something
// goes wrong with room retrieval or if the room repo return an error.
func (g *getRoom) Handle(uuid string) (*entity.Room, []*errorutil.Error) {
	room, err := g.repo.GetRoom(uuid)
	if err != nil {
		return nil, err
	}

	return room, nil
}

// NewGetRoom creates a new GetRoom.
func NewGetRoom(repo protocol.GetRoomRepo) GetRoom {
	return &getRoom{
		repo,
	}
}
