package usecase_test

import (
	"testing"

	"github.com/christian-gama/go-booking-api/internal/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/domain/error"
	"github.com/christian-gama/go-booking-api/internal/usecase"
	"github.com/christian-gama/go-booking-api/mocks"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateRoomTestSuite struct {
	suite.Suite

	createRoom usecase.CreateRoomUsecase
	repo       *mocks.SaveRoomRepo
	uuid       *mocks.UUID

	input *usecase.CreateRoomInput
}

func (s *CreateRoomTestSuite) SetupTest() {
	s.repo = mocks.NewSaveRoomRepo(s.T())
	s.uuid = mocks.NewUUID(s.T())
	s.createRoom = usecase.NewCreateRoom(s.repo, s.uuid)
	s.input = &usecase.CreateRoomInput{
		Name:        "name",
		Description: "description",
		BedCount:    1,
		Price:       1.0,
	}
}

func (s *CreateRoomTestSuite) TestNewCreateRoom_NotNil() {
	s.NotNil(s.createRoom)
}

func (s *CreateRoomTestSuite) TestCreateRoom_Handle_Success() {
	s.uuid.On("Generate").Return("uuid")
	s.repo.On("SaveRoom", mock.Anything).Return(&entity.Room{}, nil)

	result, err := s.createRoom.Handle(s.input)

	s.NotNil(result)
	s.Nil(err)
}

func (s *CreateRoomTestSuite) TestCreateRoom_Handle_InvalidInput() {
	s.uuid.On("Generate").Return("")

	result, err := s.createRoom.Handle(s.input)

	s.Nil(result)
	s.NotNil(err[0].Code, error.InvalidArgument)
}

func (s *CreateRoomTestSuite) TestCreateRoom_Handle_SaveRoomError() {
	s.uuid.On("Generate").Return("uuid")
	s.repo.On("SaveRoom", mock.Anything).Return(nil, []*error.Error{{}})

	result, err := s.createRoom.Handle(s.input)

	s.Nil(result)
	s.NotNil(err[0])
}

func TestCreateRoomTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(CreateRoomTestSuite))
}
