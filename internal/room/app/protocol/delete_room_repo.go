package protocol

import (
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
)

// DeleteRoomRepo is the interface for deleting a room.
type DeleteRoomRepo interface {
	DeleteRoom(uuid string) []*errorutil.Error
}
