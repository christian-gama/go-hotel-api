package usecase_test

import (
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/room/usecase"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/error"
	"github.com/christian-gama/go-booking-api/mocks"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type GetRoomTestSuite struct {
	suite.Suite

	getRoom usecase.GetRoomUsecase
	repo    *mocks.GetRoomRepo
}

func (s *GetRoomTestSuite) SetupTest() {
	s.repo = mocks.NewGetRoomRepo(s.T())
	s.getRoom = usecase.NewGetRoom(s.repo)
}

func (s *GetRoomTestSuite) TestNewGetRoom_NotNil() {
	s.NotNil(s.getRoom)
}

func (s *GetRoomTestSuite) TestGetRoom_Handle_Success() {
	s.repo.On("GetRoom", mock.Anything).Return(&entity.Room{}, nil)

	result, err := s.getRoom.Handle("12345678-1234-1234-1234-1234567890ab")

	s.NotNil(result)
	s.Nil(err)
}

func (s *GetRoomTestSuite) TestGetRoom_Handle_GetRoomError() {
	s.repo.On("GetRoom", mock.Anything).Return(nil, error.Add(error.New("", "", "", "")))

	result, err := s.getRoom.Handle("12345678-1234-1234-1234-1234567890ab")

	s.Nil(result)
	s.NotNil(err[0])
}

func TestGetRoomTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(GetRoomTestSuite))
}
