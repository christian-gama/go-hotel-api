package protocol

import (
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
)

// SaveRoomRepo is the interface for saving a room.
type SaveRoomRepo interface {
	SaveRoom(room *entity.Room) (*entity.Room, []*errorutil.Error)
}
