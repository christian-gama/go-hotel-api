package entity_test

import (
	"strings"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type RestrictionTestSuite struct {
	suite.Suite

	uuid        string
	name        string
	description string
}

func (s *RestrictionTestSuite) SetupTest() {
	s.uuid = "12345678-1234-1234-1234-123456789012"
	s.name = "Any name"
	s.description = strings.Repeat("a", entity.MaxRoomDescriptionLen)
}

func (s *RestrictionTestSuite) TestNewRestriction_Success() {
	result, err := entity.NewRestriction(s.uuid, s.name, s.description)

	s.NotNil(result)
	s.Nil(err)
}

func (s *RestrictionTestSuite) TestNewRestriction_UuidEmptyError() {
	result, err := entity.NewRestriction("", s.name, s.description)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("uuid", err[0].Param)
}

func (s *RestrictionTestSuite) TestNewRestriction_NameEmptyError() {
	result, err := entity.NewRestriction(s.uuid, "", s.description)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("name", err[0].Param)
}

func (s *RestrictionTestSuite) TestNewRestriction_MinDescriptionLenError() {
	description := strings.Repeat("a", entity.MinRestrictionDescriptionLen-1)

	result, err := entity.NewRestriction(s.uuid, s.name, description)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("description", err[0].Param)
}

func (s *RestrictionTestSuite) TestNewRestriction_MaxDescriptionLenError() {
	description := strings.Repeat("a", entity.MaxRestrictionDescriptionLen+1)

	result, err := entity.NewRestriction(s.uuid, s.name, description)

	s.Nil(result)
	s.Equal(errorutil.InvalidArgument, err[0].Code)
	s.Equal("description", err[0].Param)
}

func TestRestrictionTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(RestrictionTestSuite))
}
