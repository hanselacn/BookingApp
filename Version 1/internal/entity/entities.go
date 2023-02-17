package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	Name     string    `json:"name" db:"name"`
	Role     int8      `json:"role" db:"role"`
}

type Room struct {
	ID        uuid.UUID `json:"id" db:"id"`
	RoomName  string    `json:"room_name" db:"room_name"`
	CreatedBy string    `json:"created_by" db:"created_by"`
}

type Booking struct {
	ID         uuid.UUID `json:"id" db:"id"`
	BookedBy   string    `json:"booked_by" db:"booked_by"`
	BookedRoom string    `json:"booked_room" db:"booked_room"`
	BookedDay  string    `json:"booked_day" db:"booked_day"`
	Session    string    `json:"session" db:"session"`
}

type Day struct {
	ID        uuid.UUID `json:"id" db:"id"`
	DayName   string    `json:"day_name" db:"day_name"`
	CreatedBy string    `json:"created_by" db:"created_by"`
}

type Login struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	Role     int8      `json:"role" db:"role"`
	Logged   bool      `json:"logged" db:"logged"`
}
