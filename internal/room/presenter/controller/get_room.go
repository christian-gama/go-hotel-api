package controller

import (
	"github.com/christian-gama/go-booking-api/internal/room/usecase"
	"github.com/christian-gama/go-booking-api/internal/shared/presenter/controller"
	"github.com/christian-gama/go-booking-api/internal/shared/presenter/http/request"
	"github.com/christian-gama/go-booking-api/internal/shared/presenter/http/response"
)

type getRoom struct {
	getRoomUsecase usecase.GetRoomUsecase
}

// Handle is a function that handles a room's getting request.
func (g *getRoom) Handle(req *request.Request) *response.Response {
	uuid := req.Param("uuid")

	room, errs := g.getRoomUsecase.Handle(uuid)
	if errs != nil {
		return response.Exception(errs)
	}

	return response.OK(room)
}

// NewGetRoom returns a new instance of a controller that handles a room's getting.
func NewGetRoom(getRoomUsecase usecase.GetRoomUsecase) controller.Controller {
	return &getRoom{getRoomUsecase}
}
