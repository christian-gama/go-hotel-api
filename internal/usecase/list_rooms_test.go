package usecase_test

import (
	"testing"

	"github.com/christian-gama/go-booking-api/internal/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/usecase"
	"github.com/christian-gama/go-booking-api/mocks"
	"github.com/christian-gama/go-booking-api/test"
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
	s.repo.On("ListRooms").Return(nil, []*errorutil.Error{{}})

	result, err := s.listRooms.Handle()

	s.Nil(result)
	s.NotNil(err[0])
}

func TestListTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(ListRoomsTestSuite))
}
