package controller

import (
	"github.com/christian-gama/go-booking-api/internal/presenter/http/request"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/response"
	"github.com/christian-gama/go-booking-api/internal/usecase"
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
func NewCreateRoom(createRoomUsecase usecase.CreateRoomUsecase) Controller {
	return &createRoom{createRoomUsecase}
}
