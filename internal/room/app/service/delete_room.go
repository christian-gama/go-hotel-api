package service

import (
	"github.com/christian-gama/go-booking-api/internal/room/app/repo"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
)

type DeleteRoom interface {
	Handle(uuid string) (bool, []*errorutil.Error)
}

type deleteRoom struct {
	repo repo.Room
}

func (d *deleteRoom) Handle(uuid string) (bool, []*errorutil.Error) {
	if err := d.repo.DeleteRoom(uuid); err != nil {
		return false, err
	}

	return true, nil
}

func NewDeleteRoom(repo repo.Room) DeleteRoom {
	return &deleteRoom{
		repo,
	}
}
