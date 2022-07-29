package uuid

// UUID is a type that represents a universally unique identifier.
type UUID interface {
	Generate() string
}
