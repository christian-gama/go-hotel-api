package notification

import (
	"strings"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/error"
)

// Error is a notification error.
type Error struct {
	Code    error.ErrorCode
	Message string
	Param   string
}

// Notification represents the notification of a domain, which is a collection of errors.
type Notification struct {
	context string
	errors  error.Errors
}

// AddErrorf adds an error to the notification.
func (n *Notification) AddError(err *Error) {
	n.errors = append(
		n.errors,
		error.New(err.Code, err.Message, err.Param, n.context),
	)
}

// HasErrors returns true if the notification has errors.
func (n *Notification) HasErrors() bool {
	return len(n.errors) > 0
}

// Errors returns a slice of the errors of the notification.
func (n *Notification) Errors() error.Errors {
	return n.errors
}

// New creates a new notification with the given context.
func New(context string) *Notification {
	fmtContext := strings.ToLower(context[:1]) + context[1:]
	return &Notification{
		context: fmtContext,
	}
}
