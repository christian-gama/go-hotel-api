package protocol

import (
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
)

// GetRoomRepo is the interface for getting a room.
type GetRoomRepo interface {
	GetRoom(uuid string) (*entity.Room, []*errorutil.Error)
}
