package usecase

import (
	"github.com/christian-gama/go-booking-api/internal/room/app/protocol"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
)

// DeleteRoom is the interface for deleting a room.
type DeleteRoom interface {
	Handle(uuid string) (bool, []*errorutil.Error)
}

// deleteRoom is the implementation of DeleteRoom.
type deleteRoom struct {
	repo protocol.DeleteRoomRepo
}

// Handle received a uuid and deletes the room based on the uuid. Returns true if the room was deleted,
// false otherwise. It will return an error if something goes wrong with room deletion.
func (d *deleteRoom) Handle(uuid string) (bool, []*errorutil.Error) {
	if err := d.repo.DeleteRoom(uuid); err != nil {
		return false, err
	}

	return true, nil
}

// NewDeleteRoom creates a new DeleteRoom.
func NewDeleteRoom(repo protocol.DeleteRoomRepo) DeleteRoom {
	return &deleteRoom{
		repo,
	}
}
