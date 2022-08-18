package util

import (
	"os"
	"regexp"
)

// RootPath returns the root path of the project.
func RootPath() string {
	regex := regexp.MustCompile(`^(.*` + "go-booking-api" + `)`)
	workingDir, _ := os.Getwd()

	rootPath := regex.Find([]byte(workingDir))

	return string(rootPath)
}
