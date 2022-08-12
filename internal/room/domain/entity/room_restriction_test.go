package entity_test

import (
	"testing"
	"time"

	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type RoomRestrictionTestSuite struct {
	suite.Suite

	uuid        string
	room        *entity.Room
	restriction *entity.Restriction
	startDate   time.Time
	endDate     time.Time
}

func (s *RoomRestrictionTestSuite) SetupTest() {
	s.uuid = "12345678-1234-1234-1234-123456789012"
	room, err := entity.NewRoom("87654321-4321-4321-4321-123456789012",
		"Any name",
		"Any description",
		entity.MaxRoomBedCount-1,
		entity.MaxRoomPrice-1,
	)
	if err != nil {
		s.Fail("Error creating room")
	}
	s.room = room

	restriction, err := entity.NewRestriction(s.uuid, "Any name", "Any description")
	if err != nil {
		s.Fail("Error creating restriction")
	}
	s.restriction = restriction

	s.startDate = time.Now().Add(time.Hour * 1)
	s.endDate = time.Now().Add(time.Hour * 2)
}

func (s *RoomRestrictionTestSuite) TestNewRestriction_Success() {
	result, err := entity.NewRoomRestriction(s.uuid, s.room, s.restriction, s.startDate, s.endDate)

	s.NotNil(result)
	s.Nil(err)
}

func (s *RoomRestrictionTestSuite) TestNewRestriction_UuidEmptyError() {
	result, err := entity.NewRoomRestriction("", s.room, s.restriction, s.startDate, s.endDate)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("uuid", err[0].Param)
}

func (s *RoomRestrictionTestSuite) TestNewRestriction_RoomNilError() {
	result, err := entity.NewRoomRestriction(s.uuid, nil, s.restriction, s.startDate, s.endDate)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("room", err[0].Param)
}

func (s *RoomRestrictionTestSuite) TestNewRestriction_RestrictionNilError() {
	result, err := entity.NewRoomRestriction(s.uuid, s.room, nil, s.startDate, s.endDate)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("restriction", err[0].Param)
}

func (s *RoomRestrictionTestSuite) TestNewRestriction_StartDateCurrentTimeError() {
	result, err := entity.NewRoomRestriction(s.uuid, s.room, s.restriction, time.Now(), s.endDate)

	s.Nil(result)
	s.Equal(errorutil.ConditionNotMet, err[0].Code)
	s.Equal("startDate", err[0].Param)
}

func (s *RoomRestrictionTestSuite) TestNewRestriction_StartDateAfterEndTimeError() {
	result, err := entity.NewRoomRestriction(s.uuid, s.room, s.restriction, s.endDate, s.startDate)

	s.Nil(result)
	s.Equal(errorutil.ConditionNotMet, err[0].Code)
	s.Equal("startDate", err[0].Param)
}

func TestRoomRestrictionTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(RoomRestrictionTestSuite))
}
