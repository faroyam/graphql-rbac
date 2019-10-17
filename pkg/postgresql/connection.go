package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"graphql-rbac/pkg/logger"

	"go.uber.org/zap"

	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	driverName = "postgres"
)

type Connection interface {
	ExecuteInTransaction(ctx context.Context, f TransactionFunc) error
	Migrate(version uint, files []string, afn func(name string) ([]byte, error)) (uint, error)
	Close() error
}

// NewConnection returns connection to database
func NewConnection(uri string, errors map[string]error, defaultError, notFoundError error, log *logger.Logger) (Connection, error) {

	db, err := sqlx.Connect(driverName, uri)
	if err != nil {
		return nil, fmt.Errorf("connecting to db error: %w", err)
	}

	c := &connection{
		DB:     db,
		Logger: log,

		errors:        errors,
		defaultError:  defaultError,
		notFoundError: notFoundError,
	}

	return c, nil
}

type connection struct {
	*sqlx.DB
	*logger.Logger

	errors        map[string]error
	defaultError  error
	notFoundError error
}

type TransactionFunc func(tx *sqlx.Tx) error

// ExecuteInTransaction create transaction and executes the given function
func (c *connection) ExecuteInTransaction(ctx context.Context, f TransactionFunc) error {
	tx, errBegin := c.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if errBegin != nil {
		return fmt.Errorf("beginning transaction error: %w", errBegin)
	}

	errTx := f(tx)
	if errTx != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return fmt.Errorf("rollback transaction error: %w", errRollback)
		}
		return c.handle(errTx)
	}

	errCommit := tx.Commit()
	if errCommit != nil {
		return fmt.Errorf("commit transaction error: %w", errCommit)
	}

	return nil
}

func (c *connection) handle(err error) error {
	if err == nil {
		return nil
	}
	c.Debug("error caught", zap.Error(err))

	if errors.Is(err, sql.ErrNoRows) {
		return c.notFoundError
	}

	var pgErr *pq.Error
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		// unique_violation or foreign_key_violation
		case "23505", "23503":
			if err, ok := c.errors[pgErr.Constraint]; ok {
				return err
			}
		}
	}
	c.Warn("default error", zap.Error(err))
	return c.defaultError
}
