package util

import (
	"os"
	"regexp"
)

// RootPath returns the root path of the project.
func RootPath() string {
	regex := regexp.MustCompile(`^(.*` + os.Getenv("APP_NAME") + `)`)
	workingDir, _ := os.Getwd()

	rootPath := regex.Find([]byte(workingDir))

	return string(rootPath)
}
