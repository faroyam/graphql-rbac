package graphql

import (
	"context"
	"errors"
	"graphql-rbac/internal/controller"
	"graphql-rbac/internal/graphql/models"
	"graphql-rbac/internal/repository"

	"go.uber.org/zap"

	"github.com/vektah/gqlparser/gqlerror"
)

const (
	genericError    = "01"
	controllerError = "02"
	repositoryError = "03"
	modelsError     = "04"
)

// generic errors
const (
	internalError  = "01"
	directiveError = "02"
)

// controller errors
const (
	authError = "01"
)

// repository errors
const (
	notFoundError = "01"
	userError     = "02"
	roleError     = "03"
	actionError   = "04"
)

// models errors
const (
	idError = "01"
)

func gqlErrors() map[error]*gqlerror.Error {
	return map[error]*gqlerror.Error{
		// controller errors
		controller.ErrBadCredentials:    newError(controller.ErrBadCredentials.Error(), controllerError+authError+"01"),
		controller.ErrIncorrectPassword: newError(controller.ErrIncorrectPassword.Error(), controllerError+authError+"02"),
		controller.ErrInvalidToken:      newError(controller.ErrInvalidToken.Error(), controllerError+authError+"03"),
		controller.ErrTokenExpired:      newError(controller.ErrTokenExpired.Error(), controllerError+authError+"04"),

		// repository errors
		repository.ErrNotFound:               newError(repository.ErrNotFound.Error(), repositoryError+notFoundError+"01"),
		repository.ErrUserIDIsTaken:          newError(repository.ErrUserIDIsTaken.Error(), repositoryError+userError+"01"),
		repository.ErrUserLoginIsTaken:       newError(repository.ErrUserLoginIsTaken.Error(), repositoryError+userError+"02"),
		repository.ErrUserDoesNotExist:       newError(repository.ErrUserDoesNotExist.Error(), repositoryError+userError+"02"),
		repository.ErrUserAlreadyBoundToRole: newError(repository.ErrUserAlreadyBoundToRole.Error(), repositoryError+userError+"03"),
		repository.ErrRoleIDIsTaken:          newError(repository.ErrRoleIDIsTaken.Error(), repositoryError+roleError+"01"),
		repository.ErrRoleTitleIsTaken:       newError(repository.ErrRoleTitleIsTaken.Error(), repositoryError+roleError+"02"),
		repository.ErrRoleDoesNotExist:       newError(repository.ErrRoleDoesNotExist.Error(), repositoryError+roleError+"03"),
		repository.ErrActionTitleIsTaken:     newError(repository.ErrActionTitleIsTaken.Error(), repositoryError+actionError+"01"),
		repository.ErrActionDoesNotExist:     newError(repository.ErrActionDoesNotExist.Error(), repositoryError+actionError+"02"),
		// models errors
		models.ErrIncorrectIDType: newError(models.ErrIncorrectIDType.Error(), modelsError+idError+"01"),
		models.ErrIncorrectID:     newError(models.ErrIncorrectID.Error(), modelsError+idError+"02"),
	}
}

func newError(msg string, code string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: msg,
		Extensions: map[string]interface{}{
			"code": code,
		},
	}
}

func newInternalError() *gqlerror.Error {
	return &gqlerror.Error{
		Message: "internal error",
		Extensions: map[string]interface{}{
			"code": genericError + internalError + "01",
		},
	}
}

func (s *server) errorReplacer(ctx context.Context, err error) *gqlerror.Error {
	if err == nil {
		return nil
	}

	var gqlErr *gqlerror.Error
	if errors.As(err, &gqlErr) {
		return gqlErr
	}

	if customEqlErr, ok := s.errors[err]; ok {
		return customEqlErr
	}

	s.log.Error("unhandled error", zap.Error(err))

	return newInternalError()
}
