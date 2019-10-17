package resolvers

import (
	"context"
	"graphql-rbac/internal/controller"
	"graphql-rbac/internal/graphql/models"
)

func (r *queryResolver) Actions(ctx context.Context, in models.ActionsInput) ([]*models.GraphqlAction, error) {
	actions, err := r.controller.Actions(ctx, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}
	return models.ActionsFromRepo(actions)
}

func (r *mutationResolver) UpdateAction(ctx context.Context, in models.UpdateActionInput) (*models.GraphqlAction, error) {
	action, err := r.controller.UpdateAction(ctx, in.Action, in.Description)
	if err != nil {
		return nil, err
	}
	return models.ActionFromRepo(action)
}

func (r *mutationResolver) BindActionToRole(ctx context.Context, action string, roleID models.ID, binding models.Binding,
) (*models.ID, error) {
	var b controller.Binding
	switch binding.String() {
	case "BIND":
		b = controller.BIND
	default:
		b = controller.UNBIND
	}

	err := r.controller.BindActionToRole(ctx, action, roleID.UUID, b)
	return &roleID, err
}
