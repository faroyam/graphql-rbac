package server

import (
	"context"
	"fmt"
	"graphql-rbac/internal/config"
	"graphql-rbac/internal/httpctx"
	"graphql-rbac/pkg/logger"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/rs/cors"

	"go.uber.org/zap"
)

const (
	authorizationHeader  = "Authorization"
	authenticationScheme = "Bearer"
)

type Server interface {
	Run() error
	Shutdown() error
	Register(pattern string, handler http.Handler)
}

func NewServer(cfg config.GatewayConfig, log *logger.Logger) Server {
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
	})

	httpServer := &http.Server{
		Addr: cfg.Addr,
		// todo add tls config
		TLSConfig:    nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	s := &server{
		httpServer: httpServer,
		handler:    http.NewServeMux(),
		cors:       c,
		log:        log,
	}

	return s
}

type server struct {
	httpServer *http.Server
	handler    *http.ServeMux
	cors       *cors.Cors
	log        *logger.Logger
}

func (s *server) Run() error {
	s.httpServer.Handler = s.cors.Handler(s.handler)
	return s.httpServer.ListenAndServe()
}

func (s *server) Shutdown() error {
	timeout := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	s.httpServer.SetKeepAlivesEnabled(false)

	s.log.Info(fmt.Sprintf("server shutdown in %v", timeout))

	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		s.log.Error("shutdown server error")
	}

	return err
}

func (s *server) Register(pattern string, handler http.Handler) {
	s.handler.Handle(pattern, s.authMiddleware(handler))
	s.log.Info("registered handler", zap.String("pattern", pattern))
}

func (s *server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimSpace(
			strings.Replace(r.Header.Get(authorizationHeader), authenticationScheme, "", 1))
		r = r.WithContext(context.WithValue(r.Context(), httpctx.AccessTokenKey, token))
		log.Println("request")
		next.ServeHTTP(w, r)
	})
}
