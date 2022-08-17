package request_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/presenter/http/request"
	"github.com/christian-gama/go-booking-api/test"
	"github.com/stretchr/testify/suite"
)

type RequestTestSuite struct {
	suite.Suite
}

func (s *RequestTestSuite) TestRequest_ReadBody_Success() {
	body := []byte("{\"field\":\"field\",\"otherField\":\"other field\"}")
	r := &http.Request{
		Body: ioutil.NopCloser(bytes.NewReader(body)),
	}
	req := &request.Request{Request: r}

	result, err := req.ReadBody()

	s.Nil(err)
	s.NotNil(result)
}

func TestRequestTestSuite(t *testing.T) {
	test.RunUnitTest(t, new(RequestTestSuite))
}
