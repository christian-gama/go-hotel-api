package psql

import (
	"strings"

	apperror "github.com/christian-gama/go-booking-api/internal/domain/error"
)

type ErrorCode string

const (
	// ErrUniqueViolation is the error code for unique constraint violation.
	ErrUniqueViolation ErrorCode = "(SQLSTATE 23505)"

	// ErrInvalidUUID is the error code for invalid uuid.
	ErrInvalidUUID ErrorCode = "(SQLSTATE 22P02)"

	// ErrNoRows is the error code for no rows.
	ErrNoRows ErrorCode = "no rows in result set"
)

func errIs(err error, code ErrorCode) bool {
	return strings.Contains(err.Error(), string(code))
}

// Error handles the error from postgres.
func Error(err error) []*apperror.Error {
	if err == nil {
		return nil
	}

	if errIs(err, ErrUniqueViolation) {
		detail := strings.Split(err.Error(), "\"")[1]
		context := strings.Split(detail, "_")[0]
		param := strings.Split(detail, "_")[1]

		return apperror.Add(
			apperror.New(
				apperror.RepositoryError,
				"unique constraint violation",
				param,
				context,
			),
		)
	}

	if errIs(err, ErrInvalidUUID) {
		return apperror.Add(
			apperror.New(
				apperror.RepositoryError,
				"invalid uuid",
				"uuid",
				"uuid",
			),
		)
	}

	if errIs(err, ErrNoRows) {
		return apperror.Add(
			apperror.New(
				apperror.RepositoryError,
				"could not find any result",
				"rows",
				"rows",
			),
		)
	}

	return apperror.Add(
		apperror.New(
			apperror.RepositoryError,
			err.Error(),
			"repository",
			"repository",
		),
	)
}
