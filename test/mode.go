package test

import (
	"os"
	"testing"

	"github.com/christian-gama/go-booking-api/internal/shared/main/config"
	"github.com/stretchr/testify/suite"
)

func mode() string {
	config.LoadEnvFile(".env.test")
	mode := os.Getenv("TEST_MODE")

	if mode != "unit" && mode != "integration" {
		panic("TEST_MODE must be either 'unit' or 'integration'")
	}

	return mode
}

func RunIntegrationTest(t *testing.T, s suite.TestingSuite) {
	if mode() == "integration" {
		suite.Run(t, s)
	} else {
		t.Skip("Skipping integration test")
	}
}

func RunUnitTest(t *testing.T, s suite.TestingSuite) {
	if mode() == "unit" {
		suite.Run(t, s)
	} else {
		t.Skip("Skipping unit test")
	}
}
