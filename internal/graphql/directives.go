package graphql

import (
	"context"
	"graphql-rbac/internal/graphql/models"
	"graphql-rbac/internal/httpctx"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gofrs/uuid"
)

func (s *server) loginDirective(ctx context.Context, obj interface{}, next graphql.Resolver,
) (res interface{}, err error) {
	_, ok := ctx.Value(httpctx.UserIDKey).(uuid.UUID)
	if !ok {
		return nil, newError("unathorized", genericError+directiveError+"01")
	}

	return next(ctx)
}

func (s *server) authDirective(ctx context.Context, obj interface{}, next graphql.Resolver,
	action models.Action) (res interface{}, err error) {

	allowedAll, ok := ctx.Value(httpctx.AllowedAllActionsKey).(bool)
	if ok && allowedAll {
		return next(ctx)
	}

	allowedActions, ok := ctx.Value(httpctx.AllowedActionsKey).([]string)
	if ok && isInActionsList(action.String(), allowedActions) {
		return next(ctx)
	}

	return nil, newError(action.String()+" is not permitted", genericError+directiveError+"02")
}

func isInActionsList(action string, allowedActions []string) bool {
	for _, a := range allowedActions {
		if strings.EqualFold(a, action) {
			return true
		}
	}
	return false
}
