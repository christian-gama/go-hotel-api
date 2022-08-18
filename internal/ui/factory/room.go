package factory

import (
	"github.com/christian-gama/go-booking-api/internal/domain/repo"
	"github.com/christian-gama/go-booking-api/internal/infra/repo/psql"
	"github.com/christian-gama/go-booking-api/internal/presenter/controller"
	"github.com/christian-gama/go-booking-api/internal/usecase"
)

// CreateRoomRepo is a factory function that returns a new room repository.
func CreateRoomRepo() repo.SaveRoomRepo {
	conn := PsqlConn()
	dbConfig := DbConfig()

	return psql.NewRoomRepo(conn, dbConfig)
}

// CreateRoomUsecase is a factory function that returns a new room usecase.
func CreateRoomUsecase() usecase.CreateRoomUsecase {
	repo := CreateRoomRepo()
	return usecase.NewCreateRoom(repo, UUID())
}

// CreateRoomController is a factory function that returns a new room controller.
func CreateRoomController() controller.Controller {
	return controller.NewCreateRoom(CreateRoomUsecase())
}
