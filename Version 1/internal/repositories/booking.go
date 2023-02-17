package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gitlab.privy.id/go_graphql/internal/entity"
	"gitlab.privy.id/go_graphql/pkg/postgres"
)

type Booking interface {
	Create(context.Context, *entity.Booking) (*uuid.UUID, error)
	VerifyBookingBy(context.Context, string) (entity.User, error)
	VerifyBookingEntity(context.Context, string, string, string) (entity.Booking, error)
	VerifyBookingDay(context.Context, string) (entity.Day, error)
	GetBookByDayRoom(context.Context, string, string) ([]entity.Booking, error)
	GetBookByName(context.Context, string) ([]entity.Booking, error)
	Delete(context.Context, string) error
	DeleteAll(context.Context) error
}

type bookingImplementation struct {
	conn postgres.Adapter
}

func NewBookingImplementation(conn postgres.Adapter) *bookingImplementation {
	return &bookingImplementation{
		conn: conn,
	}
}

func (b bookingImplementation) Create(ctx context.Context, eb *entity.Booking) (*uuid.UUID, error) {

	queryuser := `
	SELECT 
	username
	FROM users
	WHERE username = $1
	`
	udata := b.conn.QueryRow(ctx, queryuser, eb.BookedBy)

	err := udata.Scan(
		&eb.BookedBy,
	)

	if err != nil {
		err = fmt.Errorf("error : %w", err)
		return nil, err
	}

	query := `
	INSERT INTO bookings
	(booked_by, booked_room, booked_day, session)
	VALUES ($2, $3, $4, $1) 
	RETURNING id
	`
	row := b.conn.QueryRow(ctx, query, eb.Session, eb.BookedBy, eb.BookedRoom, eb.BookedDay)

	book := entity.Booking{}

	err = row.Scan(&book.ID)

	if err != nil {
		return nil, err
	}

	return &book.ID, nil
}

func (b *bookingImplementation) VerifyBookingBy(ctx context.Context, bookedby string) (user entity.User, err error) {
	query := `SELECT id, username, password, role FROM users WHERE username = $1`

	err = b.conn.QueryRow(ctx, query, bookedby).Scan(&user.ID, &user.Username, &user.Password, &user.Role)

	if err != nil {
		err = fmt.Errorf("user doesn't exist")
		return entity.User{}, err
	}

	return user, nil
}

func (bd *bookingImplementation) VerifyBookingDay(ctx context.Context, bookedday string) (day entity.Day, err error) {
	query := `SELECT id, day_name, created_by FROM days WHERE day_name = $1`

	err = bd.conn.QueryRow(ctx, query, bookedday).Scan(&day.ID, &day.DayName, &day.CreatedBy)

	if err != nil {
		err = fmt.Errorf("day input is invalid")
		return entity.Day{}, err
	}

	return day, nil
}

func (be *bookingImplementation) VerifyBookingEntity(ctx context.Context, bookedday string, bookedses string, bookedrm string) (booking entity.Booking, err error) {
	query := `SELECT id, booked_by, booked_room, booked_day, session FROM bookings WHERE booked_day = $1 AND session = $2 AND booked_room =$3`

	err = be.conn.QueryRow(ctx, query, bookedday, bookedses, bookedrm).Scan(&booking.ID, &booking.BookedBy, &booking.BookedRoom, &booking.BookedDay, &booking.Session)

	if err != nil {
		return entity.Booking{}, err
	}
	return booking, nil
}

func (r *bookingImplementation) GetBookByDayRoom(ctx context.Context, bookedday string, bookedroom string) ([]entity.Booking, error) {
	query := `
	SELECT id, booked_by, booked_room, booked_day, session FROM bookings WHERE booked_day = $1 AND booked_room = $2
	`
	queries, err := r.conn.QueryRows(ctx, query, bookedday, bookedroom)

	if err != nil {
		err = fmt.Errorf("executing query error : %w", err)
		return nil, err
	}

	books := []entity.Booking{}

	for queries.Next() {
		var book entity.Booking

		err = queries.Scan(
			&book.ID,
			&book.BookedBy,
			&book.BookedRoom,
			&book.BookedDay,
			&book.Session,
		)

		if err != nil {
			err = fmt.Errorf("scanning bookings: %w", err)
			return nil, err
		}
		books = append(books, book)
	}

	return books, err
}

func (r *bookingImplementation) GetBookByName(ctx context.Context, bookedby string) ([]entity.Booking, error) {
	query := `
	SELECT id, booked_by, booked_room, booked_day, session FROM bookings WHERE booked_by = $1
	`
	queries, err := r.conn.QueryRows(ctx, query, bookedby)

	if err != nil {
		err = fmt.Errorf("executing query error : %w", err)
		return nil, err
	}

	books := []entity.Booking{}

	for queries.Next() {
		var book entity.Booking

		err = queries.Scan(
			&book.ID,
			&book.BookedBy,
			&book.BookedRoom,
			&book.BookedDay,
			&book.Session,
		)

		if err != nil {
			err = fmt.Errorf("scanning bookings: %w", err)
			return nil, err
		}
		books = append(books, book)
	}

	return books, err
}

func (r bookingImplementation) Delete(ctx context.Context, id string) error {
	query := `
	DELETE FROM bookings WHERE id=$1
	`

	res, err := r.conn.Exec(ctx, query, id)

	if err != nil {
		err = fmt.Errorf("executing querry error : %w", err)
		return err
	}

	deletedRow, _ := res.RowsAffected()
	if deletedRow <= 0 {
		err = fmt.Errorf("id not found")
		return err
	}
	return nil
}

func (r bookingImplementation) DeleteAll(ctx context.Context) error {
	query := `
	TRUNCATE bookings
	`

	_, err := r.conn.Exec(ctx, query)

	if err != nil {
		err = fmt.Errorf("executing querry error : %w", err)
		return err
	}
	return nil
}
