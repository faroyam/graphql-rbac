package pgrepo

import (
	"context"
	"graphql-rbac/internal/repository"
	"graphql-rbac/pkg/postgresql"
	"strings"

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
	changeSet := make(map[string]interface{}, 3)
	changeQuery := make([]string, 0, 3)

	if title != nil {
		changeQuery = append(changeQuery, "title = :title")
		changeSet["title"] = *title
	}
	if description != nil {
		changeQuery = append(changeQuery, "description = :description")
		changeSet["description"] = *description
	}

	if len(changeQuery) == 0 {
		return r.GetByID(ctx, roleID)
	}
	changeQuery = append(changeQuery, "updated_at = now()")
	subQuery := strings.Join(changeQuery, ", ")
	changeSet["id"] = roleID
	role := &repository.Role{}

	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`UPDATE roles
					SET `+subQuery+`
					WHERE id = :id
					RETURNING *;`, changeSet)
		if err != nil {
			return err
		}
		return tx.GetContext(ctx, role, query, args...)
	})
	return role, err
}

func (r *rolesRepo) Delete(ctx context.Context, roleID uuid.UUID) error {
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`DELETE FROM roles 
					WHERE id = :id;`,
			map[string]interface{}{
				"id": roleID})
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, query, args...)
		return err
	})
	return err
}

func (r *rolesRepo) Get(ctx context.Context, search *string, limit, offset int) ([]*repository.Role, error) {
	if search == nil {
		s := ""
		search = &s
	}
	roles := make([]*repository.Role, 0)

	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`SELECT * FROM roles 
					WHERE title ILIKE :search 
					OR description ILIKE :search
					LIMIT :limit
					OFFSET :offset;`,
			map[string]interface{}{
				"search": *search + "%",
				"limit":  limit,
				"offset": offset,
			})
		if err != nil {
			return err
		}

		return tx.SelectContext(ctx, &roles, query, args...)
	})
	return roles, err
}

func (r *rolesRepo) GetByID(ctx context.Context, roleID uuid.UUID) (*repository.Role, error) {
	role := &repository.Role{
		ID: roleID,
	}
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`SELECT * FROM roles 
					WHERE id = :id;`,
			role)
		if err != nil {
			return err
		}
		return tx.GetContext(ctx, role, query, args...)
	})
	return role, err
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
