package notification

import (
	"fmt"
	"strings"
)

func Error(errors []*ErrorProps) error {
	var message string
	for _, e := range errors {
		message += fmt.Sprintf("%s: %s,", e.Context, e.Message)
	}

	return fmt.Errorf(strings.Trim(message, ","))
}
