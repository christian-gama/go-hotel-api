package usecase_test

import (
	"testing"

	"github.com/christian-gama/go-hotel-api/internal/room/domain/entity"
	"github.com/christian-gama/go-hotel-api/internal/room/usecase"
	"github.com/christian-gama/go-hotel-api/internal/shared/domain/error"
	"github.com/christian-gama/go-hotel-api/mocks"
	"github.com/christian-gama/go-hotel-api/test"
	"github.com/stretchr/testify/suite"
)

type ListRoomsTestSuite struct {
	suite.Suite

	listRooms usecase.ListRoomsUsecase
	repo      *mocks.ListRoomsRepo
}

func (s *ListRoomsTestSuite) SetupTest() {
	s.repo = mocks.NewListRoomsRepo(s.T())
	s.listRooms = usecase.NewListRooms(s.repo)
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
	s.repo.On("ListRooms").Return(nil, error.Add(error.New("", "", "", "")))

	result, err := s.listRooms.Handle()

	s.Nil(result)
	s.NotNil(err[0])
}

func TestListTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(ListRoomsTestSuite))
}
