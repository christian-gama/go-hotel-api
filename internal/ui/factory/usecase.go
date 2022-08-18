package factory

import "github.com/christian-gama/go-booking-api/internal/usecase"

// ListRoomsUsecase is a factory function that returns a new room usecase.
func ListRoomsUsecase() usecase.ListRoomsUsecase {
	repo := RoomRepo()
	return usecase.NewListRooms(repo)
}

// CreateRoomUsecase is a factory function that returns a new room usecase.
func CreateRoomUsecase() usecase.CreateRoomUsecase {
	repo := RoomRepo()
	return usecase.NewCreateRoom(repo, UUID())
}

// DeleteRoomUsecase is a factory function that returns a new room usecase.
func DeleteRoomUsecase() usecase.DeleteRoomUsecase {
	repo := RoomRepo()
	return usecase.NewDeleteRoom(repo)
}
