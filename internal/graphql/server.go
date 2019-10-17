package graphql

import (
	"context"
	"graphql-rbac/internal/config"
	"graphql-rbac/internal/controller"
	"graphql-rbac/internal/graphql/resolvers"
	gqlserver "graphql-rbac/internal/graphql/server"
	"graphql-rbac/pkg/logger"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/handler"
	"github.com/vektah/gqlparser/gqlerror"
)

//go:generate go run ../../pkg/drctvgen/drctvgen.go -input=../../api/graphql/ -output=../../api/graphql/action.graphql -enum=Action -regex=(@auth.action:.)([A-Z]{2,}_?)+ -prefix=@auth.action:.
//go:generate go run github.com/99designs/gqlgen

type Server interface {
	Handler() http.Handler
}

func NewGraphqlServer(cfg config.GatewayConfig, controller *controller.Controller, log *logger.Logger) Server {
	s := &server{
		controller: controller,
		errors:     gqlErrors(),
		cfg:        cfg,
		log:        log,
	}

	return s
}

type server struct {
	controller *controller.Controller
	errors     map[error]*gqlerror.Error
	cfg        config.GatewayConfig
	log        *logger.Logger
}

func (s *server) Handler() http.Handler {
	return handler.GraphQL(s.graphqlExec(),
		handler.RequestMiddleware(s.requestMiddleware),
		handler.RecoverFunc(s.requestRecover),
		handler.ErrorPresenter(s.errorReplacer),
		handler.IntrospectionEnabled(s.cfg.EnableIntrospection),
		handler.ComplexityLimit(5),
		handler.Tracer(s.newTracer()),
		// todo caching
		// handler.EnablePersistedQueryCache()
	)
}

func (s *server) graphqlExec() graphql.ExecutableSchema {
	return gqlserver.NewExecutableSchema(gqlserver.Config{
		Resolvers: resolvers.NewResolvers(s.controller, s.log),
		Directives: gqlserver.DirectiveRoot{
			Auth:  s.authDirective,
			Login: s.loginDirective,
		},
		Complexity: gqlserver.ComplexityRoot{},
	})
}

func AddError(ctx context.Context, err error) {
	graphql.AddError(ctx, err)
}
