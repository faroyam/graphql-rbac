package pgrepo

import (
	"context"
	"graphql-rbac/internal/repository"
	"graphql-rbac/pkg/postgresql"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

func newUsersRepo(db postgresql.Connection) repository.UsersRepository {
	return &usersRepo{
		db: db,
	}
}

type usersRepo struct {
	db postgresql.Connection
}

func (r *usersRepo) Create(ctx context.Context, login, firstName, lastName, password string) (*repository.User, error) {
	user := &repository.User{
		Login:     login,
		FirstName: firstName,
		LastName:  lastName,
		Password:  password,
	}
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`INSERT INTO users (login, first_name, last_name, password)
					VALUES (:login, :first_name, :last_name, :password)
					RETURNING *;`, user)
		if err != nil {
			return err
		}
		return tx.GetContext(ctx, user, query, args...)
	})
	return user, err
}

func (r *usersRepo) Update(ctx context.Context, userID uuid.UUID, login, firstName, lastName, password *string) (*repository.User, error) {
	changeSet := make(map[string]interface{}, 5)
	changeQuery := make([]string, 0, 5)

	if login != nil {
		changeQuery = append(changeQuery, "login = :login")
		changeSet["login"] = *login
	}
	if firstName != nil {
		changeQuery = append(changeQuery, "first_name = :first_name")
		changeSet["first_name"] = *firstName
	}
	if lastName != nil {
		changeQuery = append(changeQuery, "last_name = :last_name")
		changeSet["last_name"] = *lastName
	}

	if password != nil {
		changeQuery = append(changeQuery, "password = :password")
		changeSet["password"] = *password
	}

	if len(changeQuery) == 0 {
		return r.GetByID(ctx, userID)
	}
	changeQuery = append(changeQuery, "updated_at = now()")
	subQuery := strings.Join(changeQuery, ", ")
	changeSet["id"] = userID
	user := &repository.User{}

	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`UPDATE users
					SET `+subQuery+`
					WHERE id = :id AND deleted = FALSE
					RETURNING *;`, changeSet)
		if err != nil {
			return err
		}
		return tx.GetContext(ctx, user, query, args...)
	})
	return user, err
}

func (r *usersRepo) Delete(ctx context.Context, userID uuid.UUID) error {
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`UPDATE users SET deleted = true 
					WHERE id = :id;`,
			map[string]interface{}{
				"id": userID})
		if err != nil {
			return err
		}
		_, err = tx.ExecContext(ctx, query, args...)
		return err
	})
	return err
}

func (r *usersRepo) Get(ctx context.Context, search *string, limit, offset int) ([]*repository.User, error) {
	if search == nil {
		s := ""
		search = &s
	}
	users := make([]*repository.User, 0)

	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`SELECT * FROM users 
					WHERE login ILIKE :search 
					OR first_name ILIKE :search 
                    OR last_name ILIKE :search 
					AND deleted = FALSE
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

		return tx.SelectContext(ctx, &users, query, args...)
	})
	return users, err
}

func (r *usersRepo) GetByID(ctx context.Context, userID uuid.UUID) (*repository.User, error) {
	user := &repository.User{
		ID: userID,
	}
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`SELECT * FROM users 
					WHERE id = :id AND deleted = FALSE;`, user)
		if err != nil {
			return err
		}
		return tx.GetContext(ctx, user, query, args...)
	})
	return user, err
}

func (r *usersRepo) GetByLogin(ctx context.Context, login string) (*repository.User, error) {
	user := &repository.User{
		Login: login,
	}
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`SELECT * FROM users 
					WHERE login = :login AND deleted = FALSE;`, user)
		if err != nil {
			return err
		}
		return tx.GetContext(ctx, user, query, args...)
	})
	return user, err
}

func (r *usersRepo) BindRole(ctx context.Context, userID, roleID uuid.UUID) error {
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`INSERT INTO users_roles (user_id, role_id)
					VALUES (:user_id, :role_id);`,
			map[string]interface{}{
				"user_id": userID,
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

func (r *usersRepo) UnbindRole(ctx context.Context, userID, roleID uuid.UUID) error {
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`DELETE FROM users_roles
					WHERE user_id = :user_id AND role_id = :role_id);`,
			map[string]interface{}{
				"user_id": userID,
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

func (r *usersRepo) IsSuper(ctx context.Context, userID uuid.UUID) (bool, error) {
	super := false
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`SELECT super
						FROM users_roles
							JOIN roles r on users_roles.role_id = r.id
						WHERE user_id = :user_id;`,
			map[string]interface{}{
				"user_id": userID,
			})
		if err != nil {
			return err
		}
		err = tx.GetContext(ctx, &super, query, args...)
		return err
	})
	return super, err
}

func (r *usersRepo) AllowedActions(ctx context.Context, userID uuid.UUID) ([]string, error) {
	actions := make([]string, 0)
	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`SELECT title
					FROM actions
						JOIN actions_roles mr on actions.title = mr.action_title
						JOIN users_roles ur on mr.role_id = ur.role_id
					WHERE ur.user_id = :user_id;`,
			map[string]interface{}{
				"user_id": userID,
			})
		if err != nil {
			return err
		}
		err = tx.SelectContext(ctx, &actions, query, args...)
		return err
	})
	return actions, err
}

func (r *usersRepo) Roles(ctx context.Context, userID uuid.UUID) ([]*repository.Role, error) {
	roles := make([]*repository.Role, 0)

	err := r.db.ExecuteInTransaction(ctx, func(tx *sqlx.Tx) error {
		query, args, err := tx.BindNamed(
			`SELECT *
					FROM roles
    					JOIN users_roles ur on roles.id = ur.role_id
					WHERE ur.user_id = :used_id;`,
			map[string]interface{}{
				"used_id": userID,
			})
		if err != nil {
			return err
		}

		return tx.SelectContext(ctx, &roles, query, args...)
	})
	return roles, err
}
