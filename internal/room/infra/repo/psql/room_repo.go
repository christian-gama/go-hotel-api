package psql

import (
	"context"
	"database/sql"

	"github.com/christian-gama/go-hotel-api/internal/room/domain/entity"
	"github.com/christian-gama/go-hotel-api/internal/room/domain/repo"
	"github.com/christian-gama/go-hotel-api/internal/shared/domain/error"
	"github.com/christian-gama/go-hotel-api/internal/shared/infra/config"
	"github.com/christian-gama/go-hotel-api/internal/shared/infra/repo/psql"
	"github.com/christian-gama/go-hotel-api/internal/shared/util"
)

type RoomRepo interface {
	repo.SaveRoomRepo
	repo.GetRoomRepo
	repo.ListRoomsRepo
	repo.DeleteRoomRepo
}

// roomRepoImpl is concrete implementation of the Room repository.
type roomRepoImpl struct {
	db          *sql.DB
	dbConfigger config.Db
}

// SaveRoom is the method that will save a room in the database.
func (r *roomRepoImpl) SaveRoom(room *entity.Room) (*entity.Room, error.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), r.dbConfigger.Timeout())
	defer cancel()

	stmt := `INSERT INTO room 
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
		return nil, psql.Exception(err)
	}

	return room, nil
}

// GetRoom is the method that will get a room from the database.
func (r *roomRepoImpl) GetRoom(uuid string) (*entity.Room, error.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), r.dbConfigger.Timeout())
	defer cancel()

	stmt := `SELECT uuid, name, description, bed_count, price FROM room WHERE uuid = $1`
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
		return nil, psql.Exception(err)
	}

	return room, nil
}

// ListRooms is the method that will list all the rooms from the database.
func (r *roomRepoImpl) ListRooms() ([]*entity.Room, error.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), r.dbConfigger.Timeout())
	defer cancel()

	stmt := `SELECT uuid, name, description, bed_count, price FROM room`
	rows, err := r.db.QueryContext(ctx, stmt)
	if err != nil {
		return nil, psql.Exception(err)
	}
	defer rows.Close()

	rooms := []*entity.Room{}
	for rows.Next() {
		room := &entity.Room{}
		err := rows.Scan(
			&room.UUID,
			&room.Name,
			&room.Description,
			&room.BedCount,
			&room.Price,
		)
		if err != nil {
			return nil, psql.Exception(err)
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

// DeleteRoom is the method that will delete a room from the database.
func (r *roomRepoImpl) DeleteRoom(uuid string) (bool, error.Errors) {
	ctx, cancel := context.WithTimeout(context.Background(), r.dbConfigger.Timeout())
	defer cancel()

	stmt := `DELETE FROM room WHERE uuid = $1`
	result, err := r.db.ExecContext(ctx, stmt, uuid)
	if err != nil {
		return false, psql.Exception(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, error.Add(error.New(
			error.RepositoryError,
			"could not get the number of rows affected",
			"rows",
			util.StructName(entity.Room{}),
		))
	}

	return rowsAffected > 0, nil
}

// NewRoomRepo creates a new instance of the Room repository.
func NewRoomRepo(db *sql.DB, dbConfig config.Db) RoomRepo {
	return &roomRepoImpl{
		db,
		dbConfig,
	}
}
