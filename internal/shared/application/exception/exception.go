package exception

// Error is an error type that represents a domain error.
type Error struct {
	Message string
	Name    string
}

// Exception is an interface that defines the behavior of an exception.
type Exception interface {
	BadRequest(message string) *Error
}
