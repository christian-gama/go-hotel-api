package entity

import (
	"fmt"

	"github.com/christian-gama/go-booking-api/internal/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/domain/notification"
	"github.com/christian-gama/go-booking-api/internal/util"
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
				Message: fmt.Sprintf("name cannot be longer than %d characters", MaxRestrictionNameLen),
				Param:   "name",
			},
		)
	}

	if len(r.description) > MaxRestrictionDescriptionLen {
		r.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("description cannot be longer than %d characters", MaxRestrictionDescriptionLen),
				Param:   "description",
			},
		)
	}

	if len(r.description) < MinRestrictionDescriptionLen {
		r.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: fmt.Sprintf("description cannot be shorter than %d characters", MinRestrictionDescriptionLen),
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
		notification: *notification.New(util.StructName(Restriction{})),

		UUID:        uuid,
		name:        name,
		description: description,
	}

	if err := room.validate(); err != nil {
		return nil, err
	}

	return room, nil
}
