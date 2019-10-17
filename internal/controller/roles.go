package controller

import (
	"context"
	"graphql-rbac/internal/repository"

	"github.com/gofrs/uuid"
)

func (c *Controller) Role(ctx context.Context, roleID uuid.UUID) (*repository.Role, error) {
	return c.repo.Roles().GetByID(ctx, roleID)
}

func (c *Controller) Roles(ctx context.Context, search *string, limit, offset int) ([]*repository.Role, error) {
	return c.repo.Roles().Get(ctx, search, limit, offset)
}

func (c *Controller) DeleteRole(ctx context.Context, roleID uuid.UUID) error {
	return c.repo.Roles().Delete(ctx, roleID)
}

func (c *Controller) UpdateRole(ctx context.Context, roleID uuid.UUID, title, description *string) (*repository.Role, error) {
	return c.repo.Roles().Update(ctx, roleID, title, description)
}

func (c *Controller) CreateRole(ctx context.Context, title, description string) (*repository.Role, error) {
	return c.repo.Roles().Create(ctx, title, description, false)
}

func (c *Controller) RoleActions(ctx context.Context, roleID uuid.UUID) ([]*repository.Action, error) {
	return c.repo.Roles().Actions(ctx, roleID)
}
