package factory

import (
	"github.com/christian-gama/go-booking-api/internal/domain/repo"
	"github.com/christian-gama/go-booking-api/internal/infra/repo/psql"
	"github.com/christian-gama/go-booking-api/internal/presenter/controller"
	"github.com/christian-gama/go-booking-api/internal/usecase"
)

func CreateRoomRepo() repo.SaveRoom {
	conn := PsqlConn()
	dbConfig := DbConfig()

	return psql.NewRoomRepo(conn, dbConfig)
}

func CreateRoomUsecase() usecase.CreateRoom {
	repo := CreateRoomRepo()
	return usecase.NewCreateRoom(repo, UUID())
}

func CreateRoomController() controller.Controller {
	return controller.NewCreateRoom(CreateRoomUsecase())
}
