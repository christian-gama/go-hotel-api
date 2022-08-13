package entity

import (
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
)

const (
	MaxUserPasswordLen = 32
	MinUserPasswordLen = 8
)

type User struct {
	notification *notification.Notification

	UUID     string
	Email    string
	Password string
}

func (u *User) validate() []*errorutil.Error {
	if u.UUID == "" {
		u.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		)
	}

	if u.Email == "" {
		u.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "email cannot be empty",
				Param:   "email",
			},
		)
	}

	if len(u.Password) < MinUserPasswordLen {
		u.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "password cannot be shorter than 8 characters",
				Param:   "password",
			},
		)
	}

	if len(u.Password) > MaxUserPasswordLen {
		u.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "password cannot be longer than 32 characters",
				Param:   "password",
			},
		)
	}

	if u.notification.HasErrors() {
		return u.notification.Errors()
	}

	return nil
}

func NewUser(uuid, email, password string) (*User, []*errorutil.Error) {
	user := &User{
		notification: notification.New("user"),

		UUID:     uuid,
		Email:    email,
		Password: password,
	}

	if err := user.validate(); len(err) > 0 {
		return nil, err
	}

	return user, nil
}