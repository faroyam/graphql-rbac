package resolvers

import (
	"context"
	"graphql-rbac/internal/controller"
	"graphql-rbac/internal/graphql/models"
	"graphql-rbac/internal/httpctx"
	"graphql-rbac/pkg/logger"
)

type userResolver struct {
	controller *controller.Controller
	log        *logger.Logger
}

func (r *userResolver) Roles(ctx context.Context, obj *models.User) ([]*models.Role, error) {
	roles, err := r.controller.UserRoles(ctx, obj.ID.UUID)
	if err != nil {
		return nil, err
	}
	return models.RolesFromRepo(roles)
}

func (r *queryResolver) User(ctx context.Context, id models.ID) (*models.User, error) {
	user, err := r.controller.User(ctx, id.UUID)
	if err != nil {
		return nil, err
	}
	return models.UserFromRepo(user)
}

func (r *queryResolver) Users(ctx context.Context, in models.UsersInput) ([]*models.User, error) {
	users, err := r.controller.Users(ctx, in.Search, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}
	return models.UsersFromRepo(users)
}

func (r *queryResolver) Account(ctx context.Context) (*models.User, error) {
	userID := httpctx.GetUser(ctx)

	user, err := r.controller.User(ctx, userID)
	if err != nil {
		return nil, err
	}
	return models.UserFromRepo(user)
}

func (r *mutationResolver) SignIn(ctx context.Context, in models.SignInInput) (*models.Tokens, error) {
	accessToken, refreshToken, err := r.controller.SignIn(ctx, in.Login, in.Password)
	if err != nil {
		return nil, err
	}
	return &models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (r *mutationResolver) SignUp(ctx context.Context, in models.SignUpInput) (*models.User, error) {
	newUser, err := r.controller.SignUp(ctx, in.Login, in.FirstName, in.LastName, in.Password)
	if err != nil {
		return nil, err
	}
	return models.UserFromRepo(newUser)
}

func (r *mutationResolver) SignOut(ctx context.Context) (*models.ID, error) {
	userID := httpctx.GetUser(ctx)
	err := r.controller.SignOut(ctx, userID)
	return models.IDFromUUIDPtr(userID), err
}

func (r *mutationResolver) UpdateAccount(ctx context.Context, in models.UpdateAccountInput) (*models.User, error) {
	userID := httpctx.GetUser(ctx)
	user, err := r.controller.UpdateUser(ctx, userID, in.Login, in.FirstName, in.LastName)
	if err != nil {
		return nil, err
	}
	return models.UserFromRepo(user)
}

func (r *mutationResolver) UpdateAccountPassword(ctx context.Context, in models.UpdateAccountPasswordInput) (*models.ID, error) {
	userID := httpctx.GetUser(ctx)
	err := r.controller.DeleteUser(ctx, userID)
	return models.IDFromUUIDPtr(userID), err
}

func (r *mutationResolver) DeleteAccount(ctx context.Context) (*models.ID, error) {
	userID := httpctx.GetUser(ctx)
	err := r.controller.DeleteUser(ctx, userID)
	return models.IDFromUUIDPtr(userID), err
}

func (r *mutationResolver) BindUserToRole(ctx context.Context, userID models.ID, roleID models.ID, binding models.Binding,
) (*models.ID, error) {
	var b controller.Binding
	switch binding.String() {
	case "BIND":
		b = controller.BIND
	default:
		b = controller.UNBIND
	}

	err := r.controller.BindUserToRole(ctx, userID.UUID, roleID.UUID, b)
	return &roleID, err
}
