package pgrepo

import (
	"context"
	"graphql-rbac/internal/repository"
	"graphql-rbac/pkg/postgresql"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

func newRolesRepo(db postgresql.Connection) repository.RolesRepository {
	return &rolesRepo{
		db: db,
	}
}

type rolesRepo struct {
	db postgresql.Connection
}

//todo implement functions

func (r *rolesRepo) Create(ctx context.Context, title, description string, super bool) (*repository.Role, error) {
	role := &repository.Role{
		Title:       title,
		Description: description,
		Super:       super,
	}
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`INSERT INTO roles (title, description, super)
					VALUES (:title, :description, :super)
					RETURNING *;`, role)
		if err != nil {
			return err
		}
		return tx.GetContext(ctx, role, query, args...)
	})
	return role, err
}

func (r *rolesRepo) Update(ctx context.Context, roleID uuid.UUID, title, description *string) (*repository.Role, error) {
	panic("implement me")
}

func (r *rolesRepo) Delete(ctx context.Context, roleID uuid.UUID) error {
	panic("implement me")
}

func (r *rolesRepo) Get(ctx context.Context, search *string, limit, offset int) ([]*repository.Role, error) {
	panic("implement me")
}

func (r *rolesRepo) GetByID(ctx context.Context, roleID uuid.UUID) (*repository.Role, error) {
	panic("implement me")
}

func (r *rolesRepo) GetByTitle(ctx context.Context, title string) (*repository.Role, error) {
	role := &repository.Role{
		Title: title,
	}
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`SELECT * FROM roles 
					WHERE title = :title;`, role)
		if err != nil {
			return err
		}
		return tx.GetContext(ctx, role, query, args...)
	})
	return role, err
}

func (r *rolesRepo) Actions(ctx context.Context, roleID uuid.UUID) ([]*repository.Action, error) {
	actions := make([]*repository.Action, 0)

	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`SELECT *
					FROM actions
    					JOIN actions_roles ar on actions.title = ar.action_title
					WHERE ar.role_id = :role_id;`,
			map[string]interface{}{
				"role_id": roleID,
			})
		if err != nil {
			return err
		}

		return tx.SelectContext(ctx, &actions, query, args...)
	})
	return actions, err
}
