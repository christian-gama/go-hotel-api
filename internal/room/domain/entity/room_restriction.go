package entity

import (
	"time"

	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/notification"
	"github.com/christian-gama/go-booking-api/internal/shared/util"
)

type RoomRestriction struct {
	notification *notification.Notification

	UUID        string
	Room        *Room
	Restriction *Restriction
	StartDate   time.Time
	EndDate     time.Time
}

func (r *RoomRestriction) validate() []*errorutil.Error {
	if r.UUID == "" {
		r.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "uuid cannot be empty",
				Param:   "uuid",
			},
		)
	}

	if r.Room == nil {
		r.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "room cannot be empty",
				Param:   "room",
			},
		)
	}

	if r.Restriction == nil {
		r.notification.AddError(
			&notification.Error{
				Code:    errorutil.InvalidArgument,
				Message: "restriction cannot be empty",
				Param:   "restriction",
			},
		)
	}

	if r.StartDate.After(r.EndDate) {
		r.notification.AddError(
			&notification.Error{
				Code:    errorutil.ConditionNotMet,
				Message: "start date cannot be after end date",
				Param:   "startDate",
			},
		)
	}

	if r.StartDate.Before(time.Now()) {
		r.notification.AddError(
			&notification.Error{
				Code:    errorutil.ConditionNotMet,
				Message: "start date cannot be before current time",
				Param:   "startDate",
			},
		)
	}

	if r.notification.HasErrors() {
		return r.notification.Errors()
	}

	return nil
}

func NewRoomRestriction(
	uuid string,
	room *Room,
	restriction *Restriction,
	startDate time.Time,
	endDate time.Time,
) (*RoomRestriction, []*errorutil.Error) {
	roomRestriction := &RoomRestriction{
		notification.New(util.StructName(RoomRestriction{})),

		uuid,
		room,
		restriction,
		startDate,
		endDate,
	}

	if err := roomRestriction.validate(); err != nil {
		return nil, err
	}

	return roomRestriction, nil
}
