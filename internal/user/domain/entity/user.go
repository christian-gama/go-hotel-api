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

// User represents a user with an email reference from a person entity. It hold a active status, which is always false
// by default upon creation. User is set to active once the user's email is verified. A user also holds a
// permission level that is a reference to permission entity, which determines the user's access to the system.
type User struct {
	notification *notification.Notification

	UUID            string `json:"uuid"`
	Email           string `json:"email"`
	Password        string `json:"-"`
	IsActive        bool   `json:"isActive"`
	PermissionLevel uint32 `json:"permissionLevel"`
}

// validate ensure the entity is valid. It will add an error to notification each time
// it fails a validation. It will return nil if the entity is valid.
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

// NewUser creates a new User. It will return an error if does not pass the self validation.
func NewUser(uuid, email, password string, permissionLevel uint32) (*User, error.Errors) {
	user := &User{
		notification: notification.New(util.StructName(User{})),

		UUID:            uuid,
		Email:           email,
		Password:        password,
		IsActive:        false,
		PermissionLevel: permissionLevel,
	}

	if err := user.validate(); len(err) > 0 {
		return nil, err
	}

	return user, nil
}
