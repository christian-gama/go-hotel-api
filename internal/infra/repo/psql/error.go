package psql

import (
	"fmt"
	"strings"

	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
)

type ErrorCode string

const (
	// ErrUniqueViolation is the error code for unique constraint violation.
	ErrUniqueViolation ErrorCode = "23505"

	// ErrInvalidUUID is the error code for invalid uuid.
	ErrInvalidUUID ErrorCode = "22P02"
)

func errIs(err error, code ErrorCode) bool {
	return strings.Contains(err.Error(), fmt.Sprintf("SQLSTATE %s", code))
}

// Error handles the error from postgres.
func Error(err error) []*errorutil.Error {
	if err == nil {
		return nil
	}

	if errIs(err, ErrUniqueViolation) {
		detail := strings.Split(err.Error(), "\"")[1]
		context := strings.Split(detail, "_")[0]
		param := strings.Split(detail, "_")[1]

		return []*errorutil.Error{
			{
				Code:    errorutil.RepositoryError,
				Message: "unique constraint violation",
				Context: context,
				Param:   param,
			},
		}
	}

	if errIs(err, ErrInvalidUUID) {
		return []*errorutil.Error{
			{
				Code:    errorutil.RepositoryError,
				Message: "invalid uuid",
				Context: "uuid",
				Param:   "uuid",
			},
		}
	}

	return []*errorutil.Error{
		{
			Code:    errorutil.RepositoryError,
			Message: err.Error(),
			Context: "repositoryError",
			Param:   "repositoryError",
		},
	}
}
