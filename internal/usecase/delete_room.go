package usecase

import (
	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/domain/repo"
)

type (
	// DeleteRoomUsecase is the interface for deleting a room.
	DeleteRoomUsecase interface {
		Handle(uuid string) (bool, []*errorutil.Error)
	}

	// deleteRoomImpl is the implementation of DeleteRoom.
	deleteRoomImpl struct {
		repo repo.DeleteRoomRepo
	}
)

// Handle received a uuid and deletes the room based on the uuid. Returns true if the room was deleted,
// false otherwise. It will return an error if something goes wrong with room deletion.
func (d *deleteRoomImpl) Handle(uuid string) (bool, []*errorutil.Error) {
	if err := d.repo.DeleteRoom(uuid); err != nil {
		return false, err
	}

	return true, nil
}

// NewDeleteRoom creates a new DeleteRoom.
func NewDeleteRoom(repo repo.DeleteRoomRepo) DeleteRoomUsecase {
	return &deleteRoomImpl{
		repo,
	}
}
