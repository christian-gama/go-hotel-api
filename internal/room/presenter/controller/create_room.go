package controller

import (
	"github.com/christian-gama/go-hotel-api/internal/room/usecase"
	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/controller"
	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/http/request"
	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/http/response"
)

type createRoom struct {
	createRoomUsecase usecase.CreateRoomUsecase
}

// Handle is a function that handles the room's creation request.
func (c *createRoom) Handle(req *request.Request) *response.Response {
	input := &usecase.CreateRoomInput{}
	res := response.Unmarshal(req, input)
	if res != nil {
		return res
	}

	room, errs := c.createRoomUsecase.Handle(input)
	if errs != nil {
		return response.Exception(errs)
	}

	return response.OK(room)
}

// NewCreateRoom returns a new instance of a controller that handles the room's creation.
func NewCreateRoom(createRoomUsecase usecase.CreateRoomUsecase) controller.Controller {
	return &createRoom{createRoomUsecase}
}
