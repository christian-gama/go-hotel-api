package controller

import (
	"github.com/christian-gama/go-booking-api/internal/presenter/http"
	"github.com/christian-gama/go-booking-api/internal/usecase"
)

type createRoom struct {
	createRoomUsecase usecase.CreateRoom
}

func (c *createRoom) Handle(request http.Request) http.Response {
	_, err := c.createRoomUsecase.Handle(&usecase.CreateRoomInput{})
	if err != nil {
		return http.BadRequest(err)
	}

	return http.OK([]byte("ok"))
}

func NewCreateRoom(createRoomUsecase usecase.CreateRoom) Controller {
	return &createRoom{createRoomUsecase: createRoomUsecase}
}
