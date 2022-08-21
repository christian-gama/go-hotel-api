package entity

import (
	"fmt"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/error"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
	"github.com/christian-gama/go-booking-api/internal/shared/util"
)

const (
	MaxUserPasswordLen = 32
	MinUserPasswordLen = 8
)

type User struct {
	notification *notification.Notification

	UUID     string `json:"uuid"`
	Email    string `json:"email"`
	Password string `json:"-"`
	PersonId uint32 `json:"personId"`
}

func (u *User) validate() error.Errors {
	if u.UUID == "" {
		u.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		)
	}

	if u.Email == "" {
		u.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "email cannot be empty",
				Param:   "email",
			},
		)
	}

	if len(u.Password) < MinUserPasswordLen {
		u.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: fmt.Sprintf("password cannot be shorter than %d characters", MinUserPasswordLen),
				Param:   "password",
			},
		)
	}

	if len(u.Password) > MaxUserPasswordLen {
		u.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: fmt.Sprintf("password cannot be longer than %d characters", MaxUserPasswordLen),
				Param:   "password",
			},
		)
	}

	if u.notification.HasErrors() {
		return u.notification.Errors()
	}

	return nil
}

func NewUser(uuid, email, password string, personId uint32) (*User, error.Errors) {
	user := &User{
		notification: notification.New(util.StructName(User{})),

		UUID:     uuid,
		Email:    email,
		Password: password,
		PersonId: personId,
	}

	if err := user.validate(); len(err) > 0 {
		return nil, err
	}

	return user, nil
}
