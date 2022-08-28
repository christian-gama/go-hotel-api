package controller

import (
	"github.com/christian-gama/go-hotel-api/internal/room/domain/entity"
	"github.com/christian-gama/go-hotel-api/internal/room/usecase"
	"github.com/christian-gama/go-hotel-api/internal/shared/domain/error"
	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/controller"
	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/http/request"
	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/http/response"
	"github.com/christian-gama/go-hotel-api/internal/shared/util"
)

type deleteRoom struct {
	deleteRoomUsecase usecase.DeleteRoomUsecase
}

// Handle is a function that handles the room's deletion request.
func (d *deleteRoom) Handle(req *request.Request) *response.Response {
	uuid := req.Param("uuid")

	didDelete, errs := d.deleteRoomUsecase.Handle(uuid)
	if errs != nil {
		return response.Exception(errs)
	}

	if !didDelete {
		return response.Exception(error.Add(
			error.New(
				error.NotFound,
				"could not find room with uuid",
				"uuid",
				util.StructName(entity.Room{}),
			),
		))
	}

	return response.OK(nil)
}

// NewDeleteRoom returns a new instance of a controller that handles the room's deletion.
func NewDeleteRoom(deleteRoomUsecase usecase.DeleteRoomUsecase) controller.Controller {
	return &deleteRoom{deleteRoomUsecase}
}
