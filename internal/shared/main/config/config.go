package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

// LoadEnvFile loads the .env file. It will look for the .env file in the root directory of the project.
// It will panic if cannot load the .env file.
func LoadEnvFile(envFile string) {
	err := godotenv.Load(fmt.Sprintf("%s/%s", getRootPath(), envFile))
	if err != nil {
		panic(fmt.Errorf("error on load env file %s", envFile))
	}
}

// getRootPath returns the root path of the project.
func getRootPath() string {
	regex := regexp.MustCompile(`^(.*` + "go-booking-api" + `)`)
	workingDir, _ := os.Getwd()

	rootPath := regex.Find([]byte(workingDir))
	if rootPath == nil {
		rootPath = []byte(".")
	}

	return string(rootPath)
}

// envExists checks if the environment variable exists. It will panic if does not exist.
func envExists(name string) {
	_, exists := os.LookupEnv(name)

	if !exists {
		panic(fmt.Errorf("env %s not found", name))
	}
}

// getEnv returns the environment variable. It will panic if variable is empty.
func getEnv(name string) string {
	envExists(name)

	env := strings.TrimSpace(os.Getenv(name))
	if env == "" {
		panic(fmt.Errorf("env %s is empty", name))
	}

	return env
}
