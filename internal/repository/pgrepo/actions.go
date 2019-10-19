package pgrepo

import (
	"context"
	"graphql-rbac/internal/repository"
	"graphql-rbac/pkg/postgresql"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

func newActionsRepo(db postgresql.Connection) repository.ActionsRepository {
	return &actionsRepo{
		db: db,
	}
}

type actionsRepo struct {
	db postgresql.Connection
}

func (r *actionsRepo) Create(ctx context.Context, title, description string) (*repository.Action, error) {
	action := &repository.Action{
		Title:       title,
		Description: description,
	}
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`INSERT INTO actions (title, description)
					VALUES (:title, :description)
					RETURNING *;`, action)
		if err != nil {
			return err
		}
		return tx.GetContext(ctx, action, query, args...)
	})
	return action, err
}

func (r *actionsRepo) Update(ctx context.Context, title string, description *string) (*repository.Action, error) {
	changeSet := make(map[string]interface{}, 2)
	changeQuery := make([]string, 0, 2)

	if description != nil {
		changeQuery = append(changeQuery, "description = :description")
		changeSet["description"] = *description
	}
	changeQuery = append(changeQuery, "updated_at = now()")
	subQuery := strings.Join(changeQuery, ", ")
	changeSet["title"] = title
	user := &repository.Action{}

	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`UPDATE actions
					SET `+subQuery+`
					WHERE title = :title
					RETURNING *;`, changeSet)
		if err != nil {
			return err
		}
		return tx.GetContext(ctx, user, query, args...)
	})
	return user, err
}

func (r *actionsRepo) Delete(ctx context.Context, title string) error {
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`DELETE FROM actions 
					WHERE title = :title;`,
			map[string]interface{}{
				"title": title,
			})
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, query, args...)
		return err
	})
	return err
}

func (r *actionsRepo) Get(ctx context.Context, limit, offset int) ([]*repository.Action, error) {
	actions := make([]*repository.Action, 0)
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`SELECT *
					FROM actions
					LIMIT :limit
					OFFSET :offset`,
			map[string]interface{}{
				"limit":  limit,
				"offset": offset,
			})
		if err != nil {
			return err
		}
		err = tx.SelectContext(ctx, &actions, query, args...)
		return err
	})
	return actions, err
}

func (r *actionsRepo) BindRole(ctx context.Context, action string, roleID uuid.UUID) error {

	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`INSERT INTO actions_roles (title, role_id)
					VALUES (:title, :role_id);`,
			map[string]interface{}{
				"title":   action,
				"role_id": roleID,
			})
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, query, args...)
		return err
	})
	return err
}

func (r *actionsRepo) UnbindRole(ctx context.Context, action string, roleID uuid.UUID) error {
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`DELETE FROM actions_roles
					WHERE title = :title AND role_id = :role_id);`,
			map[string]interface{}{
				"title":   action,
				"role_id": roleID,
			})
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, query, args...)
		return err
	})
	return err
}
