package request_test

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/christian-gama/go-hotel-api/internal/shared/presenter/http/request"
	"github.com/christian-gama/go-hotel-api/mocks"
	"github.com/christian-gama/go-hotel-api/test"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type RequestTestSuite struct {
	suite.Suite

	paramReader *mocks.ParamReader
}

func (s *RequestTestSuite) SetupTest() {
	s.paramReader = mocks.NewParamReader(s.T())
}

func (s *RequestTestSuite) TestRequest_ReadBody_Success() {
	body := []byte("{\"field\":\"field\",\"otherField\":\"other field\"}")
	r := &http.Request{
		Body: io.NopCloser(bytes.NewReader(body)),
	}
	req := request.New(r, nil)

	result, err := req.ReadBody()

	s.Nil(err)
	s.NotNil(result)
}

func (s *RequestTestSuite) TestRequest_Query() {
	r := &http.Request{
		Body: io.NopCloser(bytes.NewReader([]byte(""))),
		URL: &url.URL{
			RawQuery: "field=field&otherField=other_field",
		},
	}
	req := request.New(r, nil)

	result := req.Query("field")

	s.Equal("field", result)
}

func (s *RequestTestSuite) TestRequest_Param() {
	r := &http.Request{
		Body: io.NopCloser(bytes.NewReader([]byte(""))),
	}
	s.paramReader.On("Read", mock.Anything, "param").Return("value")
	req := request.New(r, s.paramReader)

	result := req.Param("param")

	s.Equal("value", result)
}

func TestRequestTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(RequestTestSuite))
}
