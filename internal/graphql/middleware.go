package graphql

import (
	"context"
	"fmt"
	"graphql-rbac/internal/controller"
	"graphql-rbac/internal/httpctx"

	"github.com/99designs/gqlgen/graphql"

	"go.uber.org/zap"
)

func (s *server) requestMiddleware(ctx context.Context, next func(ctx context.Context) []byte) []byte {
	token, ok := ctx.Value(httpctx.AccessTokenKey).(string)
	if !ok {
		return next(ctx)
	}

	userID, err := s.controller.ValidateToken(ctx, token, controller.AccessToken)
	if err != nil {
		graphql.AddError(ctx, err)
		return next(ctx)
	}
	ctx = context.WithValue(ctx, httpctx.UserIDKey, userID)

	allowedAllActions, err := s.controller.AllowedAllActions(ctx, userID)
	if err != nil {
		graphql.AddError(ctx, err)
	}
	ctx = context.WithValue(ctx, httpctx.AllowedAllActionsKey, allowedAllActions)

	allowedActions, err := s.controller.GetAllowedActions(ctx, userID)
	if err != nil {
		graphql.AddError(ctx, err)
	}
	ctx = context.WithValue(ctx, httpctx.AllowedActionsKey, allowedActions)

	return next(ctx)
}

func (s *server) requestRecover(ctx context.Context, err interface{}) error {
	s.log.Error("recovered", zap.String("error", fmt.Sprintf("%s", err)))
	return newInternalError()
}
