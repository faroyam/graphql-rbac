package repository

import (
	"time"
)

type Action struct {
	Title       string `db:"title"`
	Description string `db:"description"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
