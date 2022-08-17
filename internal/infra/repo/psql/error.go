package psql

import (
	"fmt"
	"strings"

	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
)

const (
	// ErrUniqueViolation is the error code for unique constraint violation from postgres.
	ErrUniqueViolation = "23505"
)

func errIs(err error, code string) bool {
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

	return []*errorutil.Error{
		{
			Code:    errorutil.RepositoryError,
			Message: err.Error(),
			Context: "repositoryError",
			Param:   "repositoryError",
		},
	}
}
