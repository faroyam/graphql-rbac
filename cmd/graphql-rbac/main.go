package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"graphql-rbac/internal/config"
	"graphql-rbac/internal/controller"
	"graphql-rbac/internal/graphql"
	"graphql-rbac/internal/migrations"
	"graphql-rbac/internal/repository"
	"graphql-rbac/internal/repository/pgrepo"
	"graphql-rbac/internal/server"
	"graphql-rbac/pkg/logger"
	"graphql-rbac/pkg/postgresql"
	"graphql-rbac/pkg/random"

	"github.com/99designs/gqlgen/handler"

	"go.uber.org/zap"
)

func main() {

	shutdownSignal := make(chan os.Signal, 0)
	defer func() { close(shutdownSignal) }()
	wg := sync.WaitGroup{}
	defer func() { wg.Wait() }()

	// Create logger. LOG_LEVEL environment variable used to set logging priority.
	// Higher levels are more important. Possible values [0, 6]

	log, err := logger.NewLogger(os.Getenv("LOG_LEVEL"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	log.Info("logger initialized")
	defer func() { _ = log.Sync() }()

	// read config

	cfg, err := config.NewConfig(config.NewDefaults())
	if err != nil {
		log.Fatal("can't read config", zap.Error(err))
	}
	log.Info("config initialized")

	// db and migrations

	db, err := postgresql.NewConnection(cfg.DB.PostgresURI,
		repository.RepoErrors(), repository.ErrUnknown, repository.ErrNotFound, log)
	if err != nil {
		log.Fatal("can't connect to db", zap.Error(err))
	}
	log.Info("db connection initialized")
	defer func() { _ = db.Close() }()

	if err := migrations.Migrate(db, cfg.DB.MigrationVersion, log); err != nil {
		log.Fatal("can't create migrations", zap.Error(err))
	}

	// other services

	hashGenerator := random.NewHashGenerator()
	tokenGenerator := random.NewTokenGenerator(64)
	repo := pgrepo.NewRepository(db)
	ctrl := controller.NewController(repo, hashGenerator, tokenGenerator, graphql.AddError, cfg, log)

	httpServer := server.NewServer(cfg.Gateway, log)
	graphqlServer := graphql.NewGraphqlServer(cfg.Gateway, ctrl, log)

	// seed a database

	seedDB(repo, hashGenerator, cfg.Seed, log)

	// register handlers

	if cfg.Gateway.ServePlayground {
		httpServer.Register(cfg.Gateway.PlaygroundEndpoint, handler.Playground("graphql-rbac playground", cfg.Gateway.Endpoint))
	}
	httpServer.Register(cfg.Gateway.Endpoint, graphqlServer.Handler())

	// start app

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := httpServer.Run()
		if !errors.Is(err, http.ErrServerClosed) {
			log.Error("http server error", zap.Error(err))
			return
		}
	}()
	defer func() { _ = httpServer.Shutdown() }()

	// configure shutdown conditions

	signal.Notify(shutdownSignal, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	log.Info("shutting down", zap.String("signal", func() os.Signal { return <-shutdownSignal }().String()))
}
