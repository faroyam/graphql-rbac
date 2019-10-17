package repository

import (
	"context"

	"github.com/gofrs/uuid"
)

type Repository interface {
	Users() UsersRepository
	Roles() RolesRepository
	Actions() ActionsRepository
	Sessions() SessionsRepository
}

type UsersRepository interface {
	Create(ctx context.Context, login, firstName, lastName, password string) (*User, error)
	Update(ctx context.Context, userID uuid.UUID, login, firstName, lastName, password *string) (*User, error)
	Delete(ctx context.Context, userID uuid.UUID) error
	Get(ctx context.Context, search *string, limit, offset int) ([]*User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*User, error)
	GetByLogin(ctx context.Context, login string) (*User, error)

	BindRole(ctx context.Context, userID, roleID uuid.UUID) error
	UnbindRole(ctx context.Context, userID, roleID uuid.UUID) error
	IsSuper(ctx context.Context, userID uuid.UUID) (bool, error)
	AllowedActions(ctx context.Context, userID uuid.UUID) ([]string, error)
	Roles(ctx context.Context, userID uuid.UUID) ([]*Role, error)
}

type RolesRepository interface {
	Create(ctx context.Context, title, description string, super bool) (*Role, error)
	Update(ctx context.Context, roleID uuid.UUID, title, description *string) (*Role, error)
	Delete(ctx context.Context, roleID uuid.UUID) error
	Get(ctx context.Context, search *string, limit, offset int) ([]*Role, error)
	GetByID(ctx context.Context, roleID uuid.UUID) (*Role, error)
	GetByTitle(ctx context.Context, title string) (*Role, error)
	Actions(ctx context.Context, roleID uuid.UUID) ([]*Action, error)
}

type ActionsRepository interface {
	Create(ctx context.Context, title, description string) (*Action, error)
	Update(ctx context.Context, title string, description *string) (*Action, error)
	Delete(ctx context.Context, title string) error
	Get(ctx context.Context, limit, offset int) ([]*Action, error)

	BindRole(ctx context.Context, action string, roleID uuid.UUID) error
	UnbindRole(ctx context.Context, action string, roleID uuid.UUID) error
}

type SessionsRepository interface {
	Create(ctx context.Context, userID uuid.UUID, accessToken, refreshToken string, accessTokenTTL, refreshTokenTTL int) (*Session, error)
	Update(ctx context.Context, oldRefreshToken, accessToken, refreshToken string, accessTokenTTL, refreshTokenTTL int) (*Session, error)
	DeleteByID(ctx context.Context, sessionID uuid.UUID) error
	DeleteByUserID(ctx context.Context, userID uuid.UUID) error
	GetByUserID(ctx context.Context, userID uuid.UUID) ([]*Session, error)
	GetByToken(ctx context.Context, token string) (*Session, error)
}
