package entity_test

import (
	"strings"
	"testing"

	"github.com/christian-gama/go-hotel-api/internal/room/domain/entity"
	"github.com/christian-gama/go-hotel-api/internal/shared/domain/error"
	"github.com/christian-gama/go-hotel-api/test"
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
	s.Equal(error.InvalidArgument, err[0].Code)
	s.Equal("uuid", err[0].Param)
}

func (s *RestrictionTestSuite) TestNewRestriction_NameEmptyError() {
	result, err := entity.NewRestriction(s.uuid, "", s.description)

	s.Nil(result)
	s.Equal(error.InvalidArgument, err[0].Code)
	s.Equal("name", err[0].Param)
}

func (s *RestrictionTestSuite) TestNewRestriction_MaxNameLenError() {
	name := strings.Repeat("a", entity.MaxRestrictionNameLen+1)

	result, err := entity.NewRestriction(s.uuid, name, s.description)

	s.Nil(result)
	s.Equal(error.InvalidArgument, err[0].Code)
	s.Equal("name", err[0].Param)
}

func (s *RestrictionTestSuite) TestNewRestriction_MinDescriptionLenError() {
	description := strings.Repeat("a", entity.MinRestrictionDescriptionLen-1)

	result, err := entity.NewRestriction(s.uuid, s.name, description)

	s.Nil(result)
	s.Equal(error.InvalidArgument, err[0].Code)
	s.Equal("description", err[0].Param)
}

func (s *RestrictionTestSuite) TestNewRestriction_MaxDescriptionLenError() {
	description := strings.Repeat("a", entity.MaxRestrictionDescriptionLen+1)

	result, err := entity.NewRestriction(s.uuid, s.name, description)

	s.Nil(result)
	s.Equal(error.InvalidArgument, err[0].Code)
	s.Equal("description", err[0].Param)
}

func TestRestrictionTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(RestrictionTestSuite))
}
