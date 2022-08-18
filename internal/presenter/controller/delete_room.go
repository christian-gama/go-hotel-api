package controller

import (
	"github.com/christian-gama/go-booking-api/internal/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/domain/error"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/request"
	"github.com/christian-gama/go-booking-api/internal/presenter/http/response"
	"github.com/christian-gama/go-booking-api/internal/usecase"
	"github.com/christian-gama/go-booking-api/pkg/util"
)

type deleteRoom struct {
	deleteRoomUsecase usecase.DeleteRoomUsecase
}

// Handle is a function that handles the room's deletion request.
func (d *deleteRoom) Handle(req *request.Request) *response.Response {
	uuid := req.Param("uuid")

	didDelete, errs := d.deleteRoomUsecase.Handle(uuid)
	if errs != nil {
		return response.Error(errs)
	}

	if !didDelete {
		return response.Error([]*error.Error{
			{
				Code:    error.NotFound,
				Message: "could not find room with uuid",
				Context: util.StructName(entity.Room{}),
				Param:   "uuid",
			},
		})
	}

	return response.OK(nil)
}

// NewDeleteRoom returns a new instance of a controller that handles the room's deletion.
func NewDeleteRoom(deleteRoomUsecase usecase.DeleteRoomUsecase) Controller {
	return &deleteRoom{deleteRoomUsecase}
}
