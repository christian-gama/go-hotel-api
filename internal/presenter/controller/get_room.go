package controller

import (
	"github.com/christian-gama/go-booking-api/internal/presenter/http/request"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/response"
	"github.com/christian-gama/go-booking-api/internal/usecase"
)

type getRoom struct {
	getRoomUsecase usecase.GetRoomUsecase
}

// Handle is a function that handles a room's getting request.
func (g *getRoom) Handle(req *request.Request) *response.Response {
	uuid := req.Param("uuid")

	room, errs := g.getRoomUsecase.Handle(uuid)
	if errs != nil {
		return response.Error(errs)
	}

	return response.OK(room)
}

// NewGetRoom returns a new instance of a controller that handles a room's getting.
func NewGetRoom(getRoomUsecase usecase.GetRoomUsecase) Controller {
	return &getRoom{getRoomUsecase}
}
