package service_test

import (
	"fmt"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/app/service"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/mocks"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type ListRoomsTestSuite struct {
	suite.Suite

	listRooms service.ListRooms
	repo      *mocks.Room
}

func (s *ListRoomsTestSuite) SetupTest() {
	fmt.Println("SetupTest")
	s.repo = mocks.NewRoom(s.T())
	s.listRooms = service.NewListRooms(s.repo)
}

func (s *ListRoomsTestSuite) TestNewListRooms_NotNil() {
	s.NotNil(s.listRooms)
}

func (s *ListRoomsTestSuite) TestListRooms_Handle_Success() {
	s.repo.On("ListRooms").Return([]*entity.Room{}, nil)

	result, err := s.listRooms.Handle()

	s.NotNil(result)
	s.Nil(err)
}

func (s *ListRoomsTestSuite) TestListRooms_Handle_ListRoomsError() {
	s.repo.On("ListRooms").Return(nil, []*errorutil.Error{{}})

	result, err := s.listRooms.Handle()

	s.Nil(result)
	s.NotNil(err[0])
}

func TestListTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(ListRoomsTestSuite))
}
