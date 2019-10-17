package resolvers

import (
	"context"
	"graphql-rbac/internal/controller"
	"graphql-rbac/internal/graphql/models"
	"graphql-rbac/pkg/logger"
)

type roleResolver struct {
	controller *controller.Controller
	log        *logger.Logger
}

func (r *roleResolver) Actions(ctx context.Context, obj *models.Role) ([]*models.GraphqlAction, error) {
	role, err := r.controller.RoleActions(ctx, obj.ID.UUID)
	if err != nil {
		return nil, err
	}
	return models.ActionsFromRepo(role)
}

func (r *queryResolver) Role(ctx context.Context, id models.ID) (*models.Role, error) {
	role, err := r.controller.Role(ctx, id.UUID)
	if err != nil {
		return nil, err
	}
	return models.RoleFromRepo(role)
}

func (r *queryResolver) Roles(ctx context.Context, in models.RolesInput) ([]*models.Role, error) {
	roles, err := r.controller.Roles(ctx, in.Search, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}
	return models.RolesFromRepo(roles)
}

func (r *mutationResolver) CreateRole(ctx context.Context, in models.CreateRoleInput) (*models.Role, error) {
	role, err := r.controller.CreateRole(ctx, in.Title, in.Description)
	if err != nil {
		return nil, err
	}
	return models.RoleFromRepo(role)
}

func (r *mutationResolver) UpdateRole(ctx context.Context, in models.UpdateRoleInput) (*models.Role, error) {
	role, err := r.controller.UpdateRole(ctx, in.ID.UUID, in.Title, in.Description)
	if err != nil {
		return nil, err
	}
	return models.RoleFromRepo(role)
}

func (r *mutationResolver) DeleteRole(ctx context.Context, roleID models.ID) (*models.ID, error) {
	err := r.controller.DeleteRole(ctx, roleID.UUID)
	return &roleID, err
}
