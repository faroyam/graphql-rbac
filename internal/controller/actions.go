package controller

import (
	"context"
	"graphql-rbac/internal/repository"

	"github.com/gofrs/uuid"
)

func (c *Controller) Actions(ctx context.Context, limit, offset int) ([]*repository.Action, error) {
	return c.repo.Actions().Get(ctx, limit, offset)
}

func (c *Controller) UpdateAction(ctx context.Context, action string, description *string) (*repository.Action, error) {
	return c.repo.Actions().Update(ctx, action, description)
}

func (c *Controller) BindActionToRole(ctx context.Context, action string, roleID uuid.UUID, binding Binding) error {
	var err error
	switch binding {
	case BIND:
		err = c.repo.Actions().BindRole(ctx, action, roleID)
	default:
		err = c.repo.Actions().UnbindRole(ctx, action, roleID)
	}
	return err
}
