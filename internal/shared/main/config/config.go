package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnvFile(envFile string) {
	err := godotenv.Load(fmt.Sprintf("%s/%s", getRootPath(), envFile))
	if err != nil {
		panic(fmt.Errorf("error on load env file %s", envFile))
	}
}

func getRootPath() string {
	regex := regexp.MustCompile(`^(.*` + "go-booking-api" + `)`)
	workingDir, _ := os.Getwd()

	rootPath := regex.Find([]byte(workingDir))
	if rootPath == nil {
		rootPath = []byte(".")
	}

	return string(rootPath)
}

func envExists(name string) {
	_, exists := os.LookupEnv(name)

	if !exists {
		panic(fmt.Errorf("env %s not found", name))
	}
}

func getEnv(name string) string {
	envExists(name)

	env := strings.TrimSpace(os.Getenv(name))
	if env == "" {
		panic(fmt.Errorf("env %s is empty", name))
	}

	return env
}
