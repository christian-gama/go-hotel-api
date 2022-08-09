package service_test

import (
	"fmt"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/app/dto"
	"github.com/christian-gama/go-booking-api/internal/room/app/service"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/mocks"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateRoomServiceTestSuite struct {
	suite.Suite

	createRoom service.CreateRoomService
	repo       *mocks.Room
	uuid       *mocks.UUID

	input *dto.CreateRoom
}

func (s *CreateRoomServiceTestSuite) SetupTest() {
	fmt.Println("SetupTest")
	s.repo = mocks.NewRoom(s.T())
	s.uuid = mocks.NewUUID(s.T())
	s.createRoom = service.NewCreateRoom(s.repo, s.uuid)
	s.input = &dto.CreateRoom{
		Name:        "name",
		Description: "description",
		BedCount:    1,
		Price:       1.0,
	}
}

func (s *CreateRoomServiceTestSuite) TestNewCreateRoom_NotNil() {
	s.NotNil(s.createRoom)
}

func (s *CreateRoomServiceTestSuite) TestCreateRoom_Handle_Success() {
	s.uuid.On("Generate").Return("uuid")
	s.repo.On("SaveRoom", mock.Anything).Return(&entity.Room{}, nil)

	result, err := s.createRoom.Handle(s.input)

	s.NotNil(result)
	s.Nil(err)
}

func (s *CreateRoomServiceTestSuite) TestCreateRoom_Handle_InvalidInput() {
	s.uuid.On("Generate").Return("")

	result, err := s.createRoom.Handle(s.input)

	s.Nil(result)
	s.NotNil(err[0].Code, errorutil.InvalidArgument)
}

func (s *CreateRoomServiceTestSuite) TestCreateRoom_Handle_SaveRoomError() {
	s.uuid.On("Generate").Return("uuid")
	s.repo.On("SaveRoom", mock.Anything).Return(nil, []*errorutil.Error{{}})

	result, err := s.createRoom.Handle(s.input)

	s.Nil(result)
	s.NotNil(err[0])
}

func TestCreateRoomTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(CreateRoomServiceTestSuite))
}
