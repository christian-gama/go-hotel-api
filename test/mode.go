package test

import (
	"os"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/infra/config"
	"github.com/stretchr/testify/suite"
)

const (
	unitTest        = "unit"
	integrationTest = "integration"
	bothTest        = "both"
)

func mode() string {
	config.LoadEnvFile(".env.test")
	mode := os.Getenv("TEST_MODE")

	if mode != unitTest && mode != integrationTest && mode != bothTest {
		panic("TEST_MODE must be either 'unit', 'integration' or 'both'")
	}

	return mode
}

func RunIntegrationTest(t *testing.T, s suite.TestingSuite) {
	if mode() == integrationTest || mode() == bothTest {
		suite.Run(t, s)
	} else {
		t.SkipNow()
	}
}

func RunUnitTest(t *testing.T, s suite.TestingSuite) {
	if mode() == unitTest || mode() == bothTest {
		suite.Run(t, s)
	} else {
		t.SkipNow()
	}
}
