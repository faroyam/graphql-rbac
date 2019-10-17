package repository

import (
	"time"

	"github.com/gofrs/uuid"
)

type Role struct {
	ID          uuid.UUID `db:"id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`

	Super bool `db:"super"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
