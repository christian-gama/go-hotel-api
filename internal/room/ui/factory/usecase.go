package factory

import (
	"github.com/christian-gama/go-hotel-api/internal/room/usecase"
	sharedfty "github.com/christian-gama/go-hotel-api/internal/shared/ui/factory"
)

// ListRoomsUsecase is a factory function that returns a new room usecase.
func ListRoomsUsecase() usecase.ListRoomsUsecase {
	repo := RoomRepo()
	return usecase.NewListRooms(repo)
}

// CreateRoomUsecase is a factory function that returns a new room usecase.
func CreateRoomUsecase() usecase.CreateRoomUsecase {
	repo := RoomRepo()
	return usecase.NewCreateRoom(repo, sharedfty.UUID())
}

// DeleteRoomUsecase is a factory function that returns a new room usecase.
func DeleteRoomUsecase() usecase.DeleteRoomUsecase {
	repo := RoomRepo()
	return usecase.NewDeleteRoom(repo)
}

// GetRoomUsecase is a factory function that returns a new room usecase.
func GetRoomUsecase() usecase.GetRoomUsecase {
	repo := RoomRepo()
	return usecase.NewGetRoom(repo)
}
