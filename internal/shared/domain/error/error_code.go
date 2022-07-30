package error

// ErrorCode represents the error code of a domain error.
type ErrorCode string

const (
	// Conflict is the error code when a conflict between two or more values is found.
	Conflict ErrorCode = "CONFLICT"

	// NotFound is the error code when a resource is not found.
	NotFound ErrorCode = "NOT_FOUND"

	// Internal is the error code when an internal error is found.
	InternalError ErrorCode = "INTERNAL_ERROR"

	// InvalidArgument is the error code when some argument is invalid.
	InvalidArgument ErrorCode = "INVALID_ARGUMENT"

	// Unauthorized is the error code when the user is not authorized to perform an action.
	Unauthorized ErrorCode = "UNAUTHORIZED"

	// Forbidden is the error code when the user is not allowed to perform an action.
	Forbidden ErrorCode = "FORBIDDEN"

	// Unavailable is the error code when the resource is not available.
	Unavailable ErrorCode = "UNAVAILABLE"

	// ConditionNotMet is the error code when some condition is not met.
	ConditionNotMet ErrorCode = "CONDITION_NOT_MET"
)
