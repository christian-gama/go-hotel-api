package usecase

import (
	"github.com/christian-gama/go-booking-api/internal/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/domain/repo"
)

type (
	// ListRooms is the interface that defines the retrieval of multiple rooms.
	ListRooms interface {
		Handle() ([]*entity.Room, []*errorutil.Error)
	}

	// listRoomsImpl is a concrete implementation of the ListRooms.
	listRoomsImpl struct {
		repo repo.ListRooms
	}
)

// Handle retrieves multiple room. It will return an error if something
// goes wrong with room retrieval or if the room repo return an error.
func (l *listRoomsImpl) Handle() ([]*entity.Room, []*errorutil.Error) {
	room, err := l.repo.ListRooms()
	if err != nil {
		return nil, err
	}

	return room, nil
}

// NewListRooms creates a new ListRooms.
func NewListRooms(repo repo.ListRooms) ListRooms {
	return &listRoomsImpl{
		repo,
	}
}
