package service_test

import (
	"fmt"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/app/service"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/mocks"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type GetRoomServiceTestSuite struct {
	suite.Suite

	getRoom service.GetRoomService
	repo    *mocks.Room
}

func (s *GetRoomServiceTestSuite) SetupTest() {
	fmt.Println("SetupTest")
	s.repo = mocks.NewRoom(s.T())
	s.getRoom = service.NewGetRoom(s.repo)
}

func (s *GetRoomServiceTestSuite) TestNewGetRoom_NotNil() {
	s.NotNil(s.getRoom)
}

func (s *GetRoomServiceTestSuite) TestGetRoom_Handle_Success() {
	s.repo.On("GetRoom", mock.Anything).Return(&entity.Room{}, nil)

	result, err := s.getRoom.Handle("12345678-1234-1234-1234-1234567890ab")

	s.NotNil(result)
	s.Nil(err)
}

func (s *GetRoomServiceTestSuite) TestGetRoom_Handle_GetRoomError() {
	s.repo.On("GetRoom", mock.Anything).Return(nil, []*errorutil.Error{{}})

	result, err := s.getRoom.Handle("12345678-1234-1234-1234-1234567890ab")

	s.Nil(result)
	s.NotNil(err[0])
}

func TestGetRoomTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(GetRoomServiceTestSuite))
}
