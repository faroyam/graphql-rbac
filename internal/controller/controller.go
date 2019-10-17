package controller

import (
	"context"
	"graphql-rbac/internal/config"
	"graphql-rbac/internal/repository"
	"graphql-rbac/pkg/logger"
	"graphql-rbac/pkg/random"
)

func NewController(repo repository.Repository, hashGenerator random.HashGenerator, tokenGenerator random.TokenGenerator,
	addErrFunc func(context.Context, error), cfg config.Config, log *logger.Logger) *Controller {
	return &Controller{
		repo:           repo,
		hashGenerator:  hashGenerator,
		tokenGenerator: tokenGenerator,
		addErrFunc:     addErrFunc,
		cfg:            cfg,
		log:            log,
	}
}

type Controller struct {
	repo           repository.Repository
	hashGenerator  random.HashGenerator
	tokenGenerator random.TokenGenerator
	addErrFunc     func(context.Context, error)
	cfg            config.Config
	log            *logger.Logger
}

type Binding int

const (
	BIND = iota + 1
	UNBIND
)
