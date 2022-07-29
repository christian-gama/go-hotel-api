package notification

import (
	"fmt"
	"strings"
)

// Error represents the error properties of a notification.
type Error struct {
	// Message is the error message.
	Message string

	// Context is the context of the error, e.g. "user".
	Context string
}

// Notification represents the notification of a domain, which is a collection of errors.
type Notification struct {
	context string
	errors  []*Error
}

// AddError fadds an error to the notification.
func (n *Notification) AddError(err error) {
	n.errors = append(
		n.errors,
		&Error{Message: err.Error(), Context: n.context},
	)
}

// HasErrors returns true if the notification has errors.
func (n *Notification) HasErrors() bool {
	return len(n.errors) > 0
}

// Errors returns a slice of the errors of the notification.
func (n *Notification) Errors() []*Error {
	return n.errors
}

// Error returns the error message of the notification. It will concatenate all the error messages
// with a comma and add the context to the message. For example, if the notification has two errors,
// the message will be `context:message1,context:message2`.
func (n *Notification) Error() error {
	var message string
	for _, e := range n.errors {
		message += fmt.Sprintf("%s: %s,", e.Context, e.Message)
	}

	return fmt.Errorf(strings.Trim(message, ","))
}

// New creates a new notification with the given context.
func New(context string) *Notification {
	fmtContext := strings.ToLower(context[:1]) + context[1:]

	return &Notification{
		context: fmtContext,
	}
}
