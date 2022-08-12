package entity

import (
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
)

const (
	MaxRestrictionNameLen        = 40
	MaxRestrictionDescriptionLen = 255
	MinRestrictionDescriptionLen = 10
)

type Restriction struct {
	notification notification.Notification

	UUID        string
	name        string
	description string
}

func (r *Restriction) validate() []*errorutil.Error {
	if r.UUID == "" {
		r.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		)
	}

	if r.name == "" {
		r.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "name cannot be empty",
				Param:   "name",
			},
		)
	}

	if len(r.name) > MaxRestrictionNameLen {
		r.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "name cannot be longer than 40 characters",
				Param:   "name",
			},
		)
	}

	if len(r.description) > MaxRestrictionDescriptionLen {
		r.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "description cannot be longer than 255 characters",
				Param:   "description",
			},
		)
	}

	if len(r.description) < MinRestrictionDescriptionLen {
		r.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "description cannot be shorter than 10 characters",
				Param:   "description",
			},
		)
	}

	if r.notification.HasErrors() {
		return r.notification.Errors()
	}

	return nil
}

func NewRestriction(uuid string, name string, description string) (*Restriction, []*errorutil.Error) {
	room := &Restriction{
		notification: *notification.New("restriction"),

		UUID:        uuid,
		name:        name,
		description: description,
	}

	if err := room.validate(); err != nil {
		return nil, err
	}

	return room, nil
}
