package controller

import (
	"github.com/christian-gama/go-booking-api/internal/presenter/http/request"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/response"
	"github.com/christian-gama/go-booking-api/internal/usecase"
)

type listRooms struct {
	listRoomsUsecase usecase.ListRoomsUsecase
}

// Handle is a function that handles the room's listing request.
func (l *listRooms) Handle(req *request.Request) *response.Response {
	room, errs := l.listRoomsUsecase.Handle()
	if errs != nil {
		return response.Error(errs)
	}

	return response.OK(room)
}

// NewListRooms returns a new instance of a controller that handles the room's listing.
func NewListRooms(listRoomsUsecase usecase.ListRoomsUsecase) Controller {
	return &listRooms{listRoomsUsecase}
}
