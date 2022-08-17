package controller

import (
	"github.com/christian-gama/go-booking-api/internal/presenter/http/request"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/response"
	"github.com/christian-gama/go-booking-api/internal/usecase"
)

type createRoom struct {
	createRoomUsecase usecase.CreateRoom
}

func (c *createRoom) Handle(req *request.Request) *response.Response {
	input := &usecase.CreateRoomInput{}
	res := response.Unmarshal(req, input)
	if res != nil {
		return res
	}

	room, errs := c.createRoomUsecase.Handle(input)
	if errs != nil {
		return response.Error(errs)
	}

	return response.OK(room)
}

func NewCreateRoom(createRoomUsecase usecase.CreateRoom) Controller {
	return &createRoom{createRoomUsecase}
}
