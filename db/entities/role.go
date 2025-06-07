package entities

import "time"

// Role represents a user role in the database
type Role struct {
	ID        int64     `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Code      string    `db:"code"`
}
