package protocol

import (
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
)

// ListRoomsRepo is the interface for listing all rooms.
type ListRoomsRepo interface {
	ListRooms() ([]*entity.Room, []*errorutil.Error)
}
