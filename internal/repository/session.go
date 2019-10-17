package repository

import (
	"time"

	"github.com/gofrs/uuid"
)

type Session struct {
	ID              uuid.UUID `db:"id"`
	UserID          uuid.UUID `db:"user_id"`
	AccessToken     string    `db:"access_token"`
	AccessTokenTTL  int       `db:"access_token_ttl"`
	RefreshToken    string    `db:"refresh_token"`
	RefreshTokenTTL int       `db:"refresh_token_ttl"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
