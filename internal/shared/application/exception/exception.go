package exception

type Error struct {
	Message string
	Name    string
}

type Exception interface {
	BadRequest(message string) *Error
}
