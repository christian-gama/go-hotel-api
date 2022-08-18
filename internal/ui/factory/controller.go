package factory

import (
	"github.com/christian-gama/go-booking-api/internal/presenter/controller"
)

// CreateRoomController is a factory function that returns a new room controller.
func CreateRoomController() controller.Controller {
	createRoomUsecase := CreateRoomUsecase()
	return controller.NewCreateRoom(createRoomUsecase)
}

// ListRoomsController is a factory function that returns a new room controller.
func ListRoomsController() controller.Controller {
	listRoomsUsecase := ListRoomsUsecase()
	return controller.NewListRooms(listRoomsUsecase)
}

// DeleteRoomController is a factory function that returns a new room controller.
func DeleteRoomController() controller.Controller {
	deleteRoomUsecase := DeleteRoomUsecase()
	return controller.NewDeleteRoom(deleteRoomUsecase)
}

// DeleteRoomController is a factory function that returns a new room controller.
func GetRoomController() controller.Controller {
	getRoomUsecase := GetRoomUsecase()
	return controller.NewGetRoom(getRoomUsecase)
}

// NotFoundController is a factory function that returns a generic not found controller.
func NotFoundController() controller.Controller {
	return controller.NewNotFound()
}
