package pgrepo

import (
	"context"
	"graphql-rbac/internal/repository"
	"graphql-rbac/pkg/postgresql"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

func newSessionsRepo(db postgresql.Connection) repository.SessionsRepository {
	return &sessionsRepo{
		db: db,
	}
}

type sessionsRepo struct {
	db postgresql.Connection
}

func (r *sessionsRepo) Create(ctx context.Context, userID uuid.UUID, accessToken, refreshToken string,
	accessTokenTTL, refreshTokenTTL int) (*repository.Session, error) {
	session := &repository.Session{
		UserID:          userID,
		AccessToken:     accessToken,
		AccessTokenTTL:  accessTokenTTL,
		RefreshToken:    refreshToken,
		RefreshTokenTTL: refreshTokenTTL,
	}
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`INSERT INTO sessions (user_id, access_token, access_token_ttl, refresh_token, refresh_token_ttl)
					VALUES (:user_id, :access_token, :access_token_ttl, :refresh_token, :refresh_token_ttl)
					RETURNING *;`, session)
		if err != nil {
			return err
		}
		return tx.GetContext(ctx, session, query, args...)
	})
	return session, err
}

func (r *sessionsRepo) Update(ctx context.Context, oldRefreshToken, accessToken, refreshToken string,
	accessTokenTTL, refreshTokenTTL int) (*repository.Session, error) {

	session := &repository.Session{}

	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`UPDATE sessions
					SET 
						access_token = :access_token, 
						access_token_ttl = :access_token_ttl,
						refresh_token = :refresh_token,
						refresh_token_ttl = :refresh_token_ttl,
						updated_at = now()
					WHERE refresh_token = :old_refresh_token
					RETURNING *;`,
			map[string]interface{}{
				"access_token":      accessToken,
				"access_token_ttl":  accessTokenTTL,
				"refresh_token":     refreshToken,
				"refresh_token_ttl": refreshTokenTTL,
				"old_refresh_token": oldRefreshToken,
			})
		if err != nil {
			return err
		}
		return tx.GetContext(ctx, session, query, args...)
	})
	return session, err
}

func (r *sessionsRepo) DeleteByID(ctx context.Context, sessionID uuid.UUID) error {
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`DELETE FROM sessions 
					WHERE id = :id;`,
			map[string]interface{}{
				"id": sessionID})
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, query, args...)
		return err
	})
	return err
}

func (r *sessionsRepo) DeleteByUserID(ctx context.Context, userID uuid.UUID) error {
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`DELETE FROM sessions 
					WHERE user_id = :user_id;`,
			map[string]interface{}{
				"user_id": userID})
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, query, args...)
		return err
	})
	return err
}

func (r *sessionsRepo) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*repository.Session, error) {
	sessions := make([]*repository.Session, 0)
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`SELECT * FROM sessions 
					WHERE user_id = :user_id;`,
			map[string]interface{}{
				"user_id": userID,
			})
		if err != nil {
			return err
		}
		return tx.SelectContext(ctx, sessions, query, args...)
	})
	return sessions, err
}

func (r *sessionsRepo) GetByToken(ctx context.Context, token string) (*repository.Session, error) {
	session := &repository.Session{
		AccessToken:  token,
		RefreshToken: token,
	}
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`SELECT * FROM sessions 
					WHERE access_token = :access_token OR refresh_token = :refresh_token;`,
			session)
		if err != nil {
			return err
		}
		return tx.GetContext(ctx, session, query, args...)
	})
	return session, err
}
