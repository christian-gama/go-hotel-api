package factory

import (
	"github.com/christian-gama/go-booking-api/internal/domain/repo"
	"github.com/christian-gama/go-booking-api/internal/infra/repo/psql"
	"github.com/christian-gama/go-booking-api/internal/presenter/controller"
	"github.com/christian-gama/go-booking-api/internal/usecase"
)

type roomRepo interface {
	repo.SaveRoomRepo
	repo.DeleteRoomRepo
	repo.GetRoomRepo
	repo.ListRoomsRepo
}

// RoomRepo is a factory function that returns a new room repository.
func RoomRepo() roomRepo {
	conn := PsqlConn()
	dbConfig := DbConfig()

	return psql.NewRoomRepo(conn, dbConfig)
}

// CreateRoomUsecase is a factory function that returns a new room usecase.
func CreateRoomUsecase() usecase.CreateRoomUsecase {
	repo := RoomRepo()
	return usecase.NewCreateRoom(repo, UUID())
}

// CreateRoomController is a factory function that returns a new room controller.
func CreateRoomController() controller.Controller {
	createRoomUsecase := CreateRoomUsecase()
	return controller.NewCreateRoom(createRoomUsecase)
}

// ListRoomsUsecase is a factory function that returns a new room usecase.
func ListRoomsUsecase() usecase.ListRoomsUsecase {
	repo := RoomRepo()
	return usecase.NewListRooms(repo)
}

// ListRoomsController is a factory function that returns a new room controller.
func ListRoomsController() controller.Controller {
	listRoomsUsecase := ListRoomsUsecase()
	return controller.NewListRooms(listRoomsUsecase)
}
