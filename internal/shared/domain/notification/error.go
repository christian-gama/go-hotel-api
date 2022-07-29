package notification

import (
	"fmt"
	"strings"
)

// Error returns the error message of the notification. It will concatenate all the error messages
// with a comma and add the context to the message. For example, if the notification has two errors,
// the message will be `context:message1,context:message2`.
func Error(errors []*ErrorProps) error {
	var message string
	for _, e := range errors {
		message += fmt.Sprintf("%s: %s,", e.Context, e.Message)
	}

	return fmt.Errorf(strings.Trim(message, ","))
}
