package error

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
