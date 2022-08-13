package service

import (
	"github.com/christian-gama/go-booking-api/internal/room/app/repo"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
)

// ListRooms is the interface that defines the retrieval of multiple rooms.
type ListRooms interface {
	Handle() ([]*entity.Room, []*errorutil.Error)
}

// listRooms is a concrete implementation of the ListRooms.
type listRooms struct {
	repo repo.Room
}

// Handle retrieves multiple room. It will return an error if something
// goes wrong with room retrieval or if the room repo return an error.
func (l *listRooms) Handle() ([]*entity.Room, []*errorutil.Error) {
	room, err := l.repo.ListRooms()
	if err != nil {
		return nil, err
	}

	return room, nil
}

// NewListRooms creates a new ListRooms.
func NewListRooms(repo repo.Room) ListRooms {
	return &listRooms{
		repo,
	}
}
