package service

import (
	"github.com/christian-gama/go-booking-api/internal/room/app/repo"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
)

type GetRoomService interface {
	Handle(uuid string) (*entity.Room, []*errorutil.Error)
}

type getRoom struct {
	repo repo.Room
}

func (g *getRoom) Handle(uuid string) (*entity.Room, []*errorutil.Error) {
	room, err := g.repo.GetRoom(uuid)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func NewGetRoom(repo repo.Room) GetRoomService {
	return &getRoom{
		repo,
	}
}
