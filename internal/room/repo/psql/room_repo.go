package psql

import (
	"context"
	"database/sql"

	"github.com/christian-gama/go-booking-api/internal/room/app/repo"
	"github.com/christian-gama/go-booking-api/internal/room/domain/entity"
	"github.com/christian-gama/go-booking-api/internal/shared/domain/errorutil"
	"github.com/christian-gama/go-booking-api/internal/shared/infra/configger"
)

type roomRepo struct {
	db          *sql.DB
	dbConfigger configger.Db
}

func (r *roomRepo) SaveRoom(room *entity.Room) (*entity.Room, []*errorutil.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.dbConfigger.Timeout())
	defer cancel()

	stmt := `INSERT INTO rooms 
					(uuid, name, description, bed_count, price) 
					VALUES ($1, $2, $3, $4, $5) 
					RETURNING uuid`

	_, err := r.db.ExecContext(
		ctx,
		stmt,
		room.UUID,
		room.Name,
		room.Description,
		room.BedCount,
		room.Price,
	)
	if err != nil {
		return nil, []*errorutil.Error{{
			Code:    errorutil.DatabaseError,
			Message: "Could not save a new room.",
			Context: "roomRepo",
			Param:   "SaveRoom",
		}}
	}

	return room, nil
}

func (r *roomRepo) GetRoom(uuid string) (*entity.Room, []*errorutil.Error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.dbConfigger.Timeout())
	defer cancel()

	stmt := `SELECT uuid, name, description, bed_count, price FROM rooms WHERE uuid = $1`
	row := r.db.QueryRowContext(ctx, stmt, uuid)

	room := &entity.Room{}
	err := row.Scan(
		&room.UUID,
		&room.Name,
		&room.Description,
		&room.BedCount,
		&room.Price,
	)
	if err != nil {
		return nil, []*errorutil.Error{{
			Code:    errorutil.DatabaseError,
			Message: "Could not get a room.",
			Context: "roomRepo",
			Param:   "GetRoom",
		}}
	}

	return room, nil
}

func NewRoomRepo(db *sql.DB, dbConfigger configger.Db) repo.Room {
	return &roomRepo{
		db,
		dbConfigger,
	}
}
