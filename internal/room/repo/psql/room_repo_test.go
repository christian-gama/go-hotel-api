package psql_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/room/repo/psql"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/mocks"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/golang-migrate/migrate/v4"
	"github.com/stretchr/testify/suite"
)

type RoomRepoTestSuite struct {
	suite.Suite

	m  *migrate.Migrate
	db *sql.DB
}

func (s *RoomRepoTestSuite) SetupSuite() {
	s.db = test.SetupDb()
	s.m = test.Migrate(s.db)
}

func (s *RoomRepoTestSuite) SetupTest() {
	err := s.m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
}

func (s *RoomRepoTestSuite) TearDownTest() {
	err := s.m.Down()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}
}

func (s *RoomRepoTestSuite) TearDownSuite() {
	s.m.Close()
}

func (s *RoomRepoTestSuite) TestRoomRepo_SaveRoom_Success() {
	room, _ := entity.NewRoom(
		"12345678-1234-1234-1234-123456789012",
		"Test Room",
		"This is a test room",
		1,
		1,
	)
	dbConfigMock := mocks.NewDb(s.T())
	roomRepo := psql.NewRoomRepo(s.db, dbConfigMock)
	dbConfigMock.On("Timeout").Return(2 * time.Second)

	_, err := roomRepo.SaveRoom(room)

	s.Nil(err)
}

func (s *RoomRepoTestSuite) TestRoomRepo_SaveRoom_Error() {
	room, _ := entity.NewRoom(
		"12345678-1234-1234-1234-123456789012",
		"Test Room",
		"This is a test room",
		1,
		1,
	)
	dbConfigMock := mocks.NewDb(s.T())
	roomRepo := psql.NewRoomRepo(s.db, dbConfigMock)
	dbConfigMock.On("Timeout").Return(1 * time.Microsecond)

	_, err := roomRepo.SaveRoom(room)

	s.Equal(errorutil.DatabaseError, err[0].Code)
}

func (s *RoomRepoTestSuite) TestRoomRepo_GetRoom_Success() {
	room, _ := entity.NewRoom(
		"12345678-1234-1234-1234-123456789012",
		"Test Room",
		"This is a test room",
		1,
		1,
	)
	dbConfigMock := mocks.NewDb(s.T())
	roomRepo := psql.NewRoomRepo(s.db, dbConfigMock)
	dbConfigMock.On("Timeout").Return(2 * time.Second)
	roomRepo.SaveRoom(room)

	result, err := roomRepo.GetRoom(room.UUID)

	s.Nil(err)
	s.Equal(room.UUID, result.UUID)
}

func (s *RoomRepoTestSuite) TestRoomRepo_GetRoom_Error() {
	room, _ := entity.NewRoom(
		"12345678-1234-1234-1234-123456789012",
		"Test Room",
		"This is a test room",
		1,
		1,
	)
	dbConfigMock := mocks.NewDb(s.T())
	roomRepo := psql.NewRoomRepo(s.db, dbConfigMock)
	mock := dbConfigMock.On("Timeout").Return(2 * time.Second)
	roomRepo.SaveRoom(room)
	mock.Unset()
	mock.On("Timeout").Return(1 * time.Microsecond)

	_, err := roomRepo.GetRoom(room.UUID)

	s.Equal(errorutil.DatabaseError, err[0].Code)
}

func TestRoomRepoTestSuite(t *testing.T) {
	test.RunIntegrationTest(t, new(RoomRepoTestSuite))
}
