package controller_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/christian-gama/go-hotel-api/internal/shared/domain/error"
	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/controller"
	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/http/request"
	"github.com/christian-gama/go-hotel-api/mocks"
	"github.com/christian-gama/go-hotel-api/test"
	"github.com/stretchr/testify/suite"
)

type NotFoundTestSuite struct {
	suite.Suite

	ctrl        controller.Controller
	paramReader *mocks.ParamReader
}

func (s *NotFoundTestSuite) SetupTest() {
	s.paramReader = mocks.NewParamReader(s.T())
	s.ctrl = controller.NewNotFound()
}

func (s *NotFoundTestSuite) TestNewNotFound_NotNil() {
	s.NotNil(s.ctrl)
}

func (s *NotFoundTestSuite) TestNotFound_Handle() {
	req := request.New(httptest.NewRequest(http.MethodGet, "/any_url", nil), s.paramReader)

	result := s.ctrl.Handle(req)

	s.Equal(error.NotFound, result.Errors[0].Code)
	s.Equal("/any_url", result.Errors[0].Param)
}

func TestNotFoundTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(NotFoundTestSuite))
}
