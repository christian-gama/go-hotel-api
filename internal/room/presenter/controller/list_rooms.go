package controller

import (
	"github.com/christian-gama/go-hotel-api/internal/room/usecase"
	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/controller"
	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/http/request"
	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/http/response"
)

type listRooms struct {
	listRoomsUsecase usecase.ListRoomsUsecase
}

// Handle is a function that handles the room's listing request.
func (l *listRooms) Handle(req *request.Request) *response.Response {
	room, errs := l.listRoomsUsecase.Handle()
	if errs != nil {
		return response.Exception(errs)
	}

	return response.OK(room)
}

// NewListRooms returns a new instance of a controller that handles the room's listing.
func NewListRooms(listRoomsUsecase usecase.ListRoomsUsecase) controller.Controller {
	return &listRooms{listRoomsUsecase}
}
