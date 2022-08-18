package factory

import (
	"github.com/christian-gama/go-booking-api/internal/room/presenter/controller"
	sharedctrl "github.com/christian-gama/go-booking-api/internal/shared/presenter/controller"
)

// CreateRoomController is a factory function that returns a new room controller.
func CreateRoomController() sharedctrl.Controller {
	createRoomUsecase := CreateRoomUsecase()
	return controller.NewCreateRoom(createRoomUsecase)
}

// ListRoomsController is a factory function that returns a new room controller.
func ListRoomsController() sharedctrl.Controller {
	listRoomsUsecase := ListRoomsUsecase()
	return controller.NewListRooms(listRoomsUsecase)
}

// DeleteRoomController is a factory function that returns a new room controller.
func DeleteRoomController() sharedctrl.Controller {
	deleteRoomUsecase := DeleteRoomUsecase()
	return controller.NewDeleteRoom(deleteRoomUsecase)
}

// DeleteRoomController is a factory function that returns a new room controller.
func GetRoomController() sharedctrl.Controller {
	getRoomUsecase := GetRoomUsecase()
	return controller.NewGetRoom(getRoomUsecase)
}
