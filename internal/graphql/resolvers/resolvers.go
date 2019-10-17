package resolvers

import (
	"graphql-rbac/internal/controller"
	"graphql-rbac/internal/graphql/server"
	"graphql-rbac/pkg/logger"
)

func NewResolvers(controller *controller.Controller, log *logger.Logger) server.ResolverRoot {
	return &resolverRoot{
		MutationResolver: &mutationResolver{
			controller: controller,
			log:        log,
		},
		QueryResolver: &queryResolver{
			controller: controller,
			log:        log,
		},
		UserResolver: &userResolver{
			controller: controller,
			log:        log,
		},
		RoleResolver: &roleResolver{
			controller: controller,
			log:        log,
		},
	}
}

type resolverRoot struct {
	server.MutationResolver
	server.QueryResolver

	server.UserResolver
	server.RoleResolver
}

func (r *resolverRoot) Mutation() server.MutationResolver {
	return r.MutationResolver
}

func (r *resolverRoot) Query() server.QueryResolver {
	return r.QueryResolver
}

func (r *resolverRoot) User() server.UserResolver {
	return r.UserResolver
}

func (r *resolverRoot) Role() server.RoleResolver {
	return r.RoleResolver
}

type mutationResolver struct {
	controller *controller.Controller
	log        *logger.Logger
}

type queryResolver struct {
	controller *controller.Controller
	log        *logger.Logger
}
