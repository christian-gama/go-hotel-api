package uuid_test

import (
	"testing"

	"github.com/christian-gama/go-booking-api/internal/shared/app/uuid"
	uuidImpl "github.com/christian-gama/go-booking-api/internal/shared/infra/uuid"
	"github.com/stretchr/testify/suite"
)

type UuidTestSuite struct {
	suite.Suite

	uuid uuid.UUID
}

func (s *UuidTestSuite) SetupTest() {
	s.uuid = uuidImpl.NewUUID()
}

func (s *UuidTestSuite) TestUuid_Generate() {
	result := s.uuid.Generate()
	s.Regexp("[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}", result)
}

func (s *UuidTestSuite) TestNewUuid() {
	s.NotNil(s.uuid)
}

func TestUuidTestSuite(t *testing.T) {
	suite.Run(t, new(UuidTestSuite))
}
