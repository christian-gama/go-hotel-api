package entity_test

import (
	"strings"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type RoomTestSuite struct {
	suite.Suite

	room        *entity.Room
	uuid        string
	name        string
	description string
	bedCount    uint8
	price       float32
	isAvailable bool
}

func (s *RoomTestSuite) SetupTest() {
	s.uuid = "12345678-1234-1234-1234-123456789012"
	s.name = "Any name"
	s.description = strings.Repeat("a", entity.MaxRoomDescriptionLen)
	s.bedCount = entity.MinRoomBedCount
	s.price = entity.MinRoomPrice
	s.isAvailable = false

	room, err := entity.NewRoom(s.uuid, s.name, s.description, s.bedCount, s.price, s.isAvailable)
	if err != nil {
		s.Fail("could not create a new room in the test suite")
	}

	s.room = room
}

func (s *RoomTestSuite) TestNewRoom_Success() {
	result, err := entity.NewRoom(s.uuid, s.name, s.description, s.bedCount, s.price, s.isAvailable)

	s.NotNil(result)
	s.Nil(err)
}

func (s *RoomTestSuite) TestNewRoom_UuidEmptyError() {
	result, err := entity.NewRoom("", s.name, s.description, s.bedCount, s.price, s.isAvailable)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *RoomTestSuite) TestNewRoom_NameEmptyError() {
	result, err := entity.NewRoom(s.uuid, "", s.description, s.bedCount, s.price, s.isAvailable)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *RoomTestSuite) TestNewRoom_MinDescriptionLenError() {
	description := strings.Repeat("a", entity.MinRoomDescriptionLen-1)

	result, err := entity.NewRoom(s.uuid, s.name, description, s.bedCount, s.price, s.isAvailable)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *RoomTestSuite) TestNewRoom_MaxDescriptionLenError() {
	description := strings.Repeat("a", entity.MaxRoomDescriptionLen+1)

	result, err := entity.NewRoom(s.uuid, s.name, description, s.bedCount, s.price, s.isAvailable)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *RoomTestSuite) TestNewRoom_MinBedCountError() {
	bedCount := entity.MinRoomBedCount - 1

	result, err := entity.NewRoom(s.uuid, s.name, s.description, bedCount, s.price, s.isAvailable)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *RoomTestSuite) TestNewRoom_MaxBedCountError() {
	bedCount := entity.MaxRoomBedCount + 1

	result, err := entity.NewRoom(s.uuid, s.name, s.description, bedCount, s.price, s.isAvailable)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *RoomTestSuite) TestNewRoom_MinPriceError() {
	price := entity.MinRoomPrice - 1

	result, err := entity.NewRoom(s.uuid, s.name, s.description, s.bedCount, price, s.isAvailable)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func (s *RoomTestSuite) TestNewRoom_MaxPriceError() {
	price := entity.MaxRoomPrice + 1

	result, err := entity.NewRoom(s.uuid, s.name, s.description, s.bedCount, price, s.isAvailable)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
}

func TestRoomTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(RoomTestSuite))
}
