package notification

import (
	"fmt"
	"strings"
)

// ErrorProps represents the error properties of a notification.
type ErrorProps struct {
	// Message is the error message.
	Message string

	// Context is the context of the error, e.g. "user".
	Context string
}

// Notification represents the notification of a domain, which is a collection of errors.
type Notification struct {
	context string
	errors  []*ErrorProps
}

// AddErrorf adds an error to the notification and formats the message with the given arguments.
// Is equivalent to call AddError with the result of fmt.Sprintf.
func (n *Notification) AddErrorf(message string, v ...any) {
	n.errors = append(
		n.errors,
		&ErrorProps{Message: fmt.Sprintf(message, v...), Context: n.context},
	)
}

// HasErrors returns true if the notification has errors.
func (n *Notification) HasErrors() bool {
	return len(n.errors) > 0
}

// Errors returns a slice of the errors of the notification.
func (n *Notification) Errors() []*ErrorProps {
	return n.errors
}

// New creates a new notification with the given context.
func New(context string) *Notification {
	fmtContext := strings.ToLower(context[:1]) + context[1:]

	return &Notification{
		context: fmtContext,
	}
}
