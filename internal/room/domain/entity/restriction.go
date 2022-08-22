package entity

import (
	"fmt"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/error"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
	"github.com/christian-gama/go-booking-api/internal/shared/util"
)

const (
	MaxRestrictionNameLen        = 40
	MaxRestrictionDescriptionLen = 255
	MinRestrictionDescriptionLen = 10
)

// Restriction represents a possible restriction for a room.
type Restriction struct {
	notification notification.Notification

	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (r *Restriction) validate() error.Errors {
	if r.UUID == "" {
		r.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		)
	}

	if r.Name == "" {
		r.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: "name cannot be empty",
				Param:   "name",
			},
		)
	}

	if len(r.Name) > MaxRestrictionNameLen {
		r.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: fmt.Sprintf("name cannot be longer than %d characters", MaxRestrictionNameLen),
				Param:   "name",
			},
		)
	}

	if len(r.Description) > MaxRestrictionDescriptionLen {
		r.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
				Message: fmt.Sprintf("description cannot be longer than %d characters", MaxRestrictionDescriptionLen),
				Param:   "description",
			},
		)
	}

	if len(r.Description) < MinRestrictionDescriptionLen {
		r.notification.AddError(
			&notification.Error{
				Code:    error.InvalidArgument,
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

// NewRestriction creates a new Restriction. It will return an error if does not pass the self validation.
func NewRestriction(uuid string, name string, description string) (*Restriction, error.Errors) {
	room := &Restriction{
		notification: *notification.New(util.StructName(Restriction{})),

		UUID:        uuid,
		Name:        name,
		Description: description,
	}

	if err := room.validate(); err != nil {
		return nil, err
	}

	return room, nil
}
