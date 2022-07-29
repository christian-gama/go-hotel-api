package notification

import (
	"fmt"
	"strings"
)

type ErrorProps struct {
	Message string
	Context string
}

type Notification struct {
	context string
	errors  []*ErrorProps
}

func (n *Notification) AddErrorf(message string, v ...any) {
	n.errors = append(
		n.errors,
		&ErrorProps{Message: fmt.Sprintf(message, v...), Context: n.context},
	)
}

func (n *Notification) ClearErrors() {
	n.errors = nil
}

func (n *Notification) HasErrors() bool {
	return len(n.errors) > 0
}

func (n *Notification) Errors() []*ErrorProps {
	return n.errors
}

func New(context string) *Notification {
	fmtContext := strings.ToLower(context[:1]) + context[1:]

	return &Notification{
		context: fmtContext,
	}
}
