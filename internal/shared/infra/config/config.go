package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/christian-gama/go-hotel-api/internal/shared/util"
	"github.com/joho/godotenv"
)

// LoadEnvFile loads the .env file. It will look for the .env file in the root directory of the project.
// It will panic if cannot load the .env file.
func LoadEnvFile(envFile string) {
	err := godotenv.Load(fmt.Sprintf("%s/%s", util.RootPath(), envFile))
	if err != nil {
		panic(fmt.Errorf("error on load env file %s", envFile))
	}

	preCheckEnvs()
}

// preCheckEnvs checks each environment variable. It will panic if any of them is empty
// or does not exists.
func preCheckEnvs() {
	envExists(dbHost)
	envExists(dbPort)
	envExists(dbUser)
	envExists(dbPassword)
	envExists(dbName)
	envExists(dbSgbd)
	envExists(dbSslMode)
	envExists(dbMaxConnections)
	envExists(dbMaxIdleConnections)
	envExists(dbMaxLifeTimeMin)
	envExists(dbTimeoutSec)

	envExists(env)
	envExists(appHost)
	envExists(appPort)
}

// getEnv returns the environment variable. It will panic if variable is empty.
func getEnv(name string) string {
	return strings.TrimSpace(os.Getenv(name))
}

// envExists checks if the environment variable exists and it is not empty. It will panic if does not
// exists or is empty.
func envExists(name string) {
	_, exists := os.LookupEnv(name)

	if !exists {
		panic(fmt.Errorf("env %s not found", name))
	}

	if getEnv(name) == "" {
		panic(fmt.Errorf("env %s is empty", name))
	}
}
