package errorutil

// Error is a struct that represents an error.
type Error struct {
	// Code is the error code.
	Code ErrorCode

	// Context is usually the name of the domain/entity where the error occurred.
	Context string

	// Message is the error message.
	Message string

	// Param is the name of the parameter that caused the error.
	Param string
}

// New creates a new instance of Error.
func New(code ErrorCode, message, param, context string) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Param:   param,
		Context: context,
	}
}

// Append appends multiple errors into an array.
func Append(err ...*Error) []*Error {
	return err
}
