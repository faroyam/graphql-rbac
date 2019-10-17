package pgrepo

import (
	"graphql-rbac/internal/repository"
	"graphql-rbac/pkg/postgresql"
)

func NewRepository(db postgresql.Connection) repository.Repository {
	return &repo{
		UsersRepository:    newUsersRepo(db),
		RolesRepository:    newRolesRepo(db),
		ActionsRepository:  newActionsRepo(db),
		SessionsRepository: newSessionsRepo(db),
	}
}

type repo struct {
	repository.UsersRepository
	repository.RolesRepository
	repository.ActionsRepository
	repository.SessionsRepository
}

func (r *repo) Users() repository.UsersRepository {
	return r.UsersRepository
}

func (r *repo) Roles() repository.RolesRepository {
	return r.RolesRepository
}

func (r *repo) Actions() repository.ActionsRepository {
	return r.ActionsRepository
}

func (r *repo) Sessions() repository.SessionsRepository {
	return r.SessionsRepository
}
