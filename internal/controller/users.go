package controller

import (
	"context"
	"errors"
	"graphql-rbac/internal/repository"
	"time"

	"github.com/gofrs/uuid"
)

type TokenType int

const (
	AccessToken = iota + 1
	RefreshToken
)

func (c *Controller) SignIn(ctx context.Context, login, password string) (string, string, error) {
	user, err := c.repo.Users().GetByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return "", "", ErrBadCredentials
		}
		return "", "", err
	}

	validPassword, err := c.hashGenerator.Compare(password, user.Password)
	if err != nil {
		return "", "", err
	}
	if !validPassword {
		return "", "", ErrBadCredentials
	}

	accessToken, err := c.tokenGenerator.Generate()
	if err != nil {
		return "", "", err
	}

	refreshToken, err := c.tokenGenerator.Generate()
	if err != nil {
		return "", "", err
	}

	session, err := c.repo.Sessions().Create(ctx, user.ID, accessToken, refreshToken,
		c.cfg.Session.AccessTokenTTL, c.cfg.Session.RefreshTokenTTL)
	if err != nil {
		return "", "", err
	}

	return session.AccessToken, session.RefreshToken, nil
}

func (c *Controller) SignUp(ctx context.Context, login, firstName, lastName, password string) (*repository.User, error) {

	passwordHash, err := c.hashGenerator.Generate(password)
	if err != nil {
		return nil, err
	}

	newUser, err := c.repo.Users().Create(ctx, login, firstName, lastName, passwordHash)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (c *Controller) RefreshTokens(ctx context.Context, oldRefreshToken string) (string, string, error) {

	_, err := c.validateToken(ctx, oldRefreshToken, RefreshToken)
	if err != nil {
		return "", "", err
	}

	accessToken, err := c.tokenGenerator.Generate()
	if err != nil {
		return "", "", err
	}

	refreshToken, err := c.tokenGenerator.Generate()
	if err != nil {
		return "", "", err
	}

	session, err := c.repo.Sessions().Update(ctx, oldRefreshToken, accessToken, refreshToken,
		c.cfg.Session.AccessTokenTTL, c.cfg.Session.RefreshTokenTTL)
	if err != nil {
		return "", "", err
	}

	return session.AccessToken, session.RefreshToken, nil
}

func (c *Controller) SignOut(ctx context.Context, userID uuid.UUID) error {
	return c.repo.Sessions().DeleteByUserID(ctx, userID)
}

func (c *Controller) ValidateToken(ctx context.Context, token string, tokenType TokenType) (uuid.UUID, error) {
	return c.validateToken(ctx, token, tokenType)
}

func (c *Controller) AllowedAllActions(ctx context.Context, userID uuid.UUID) (bool, error) {
	allowedAll, err := c.repo.Users().IsSuper(ctx, userID)
	if err != nil {
		return false, err
	}
	return allowedAll, nil
}

func (c *Controller) GetAllowedActions(ctx context.Context, userID uuid.UUID) ([]string, error) {
	actions, err := c.repo.Users().AllowedActions(ctx, userID)
	if err != nil {
		return nil, err
	}

	return actions, nil
}

func (c *Controller) User(ctx context.Context, userID uuid.UUID) (*repository.User, error) {
	return c.repo.Users().GetByID(ctx, userID)
}

func (c *Controller) Users(ctx context.Context, search *string, limit, offset int) ([]*repository.User, error) {
	return c.repo.Users().Get(ctx, search, limit, offset)
}

func (c *Controller) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	return c.repo.Users().Delete(ctx, userID)
}

func (c *Controller) UpdateUser(ctx context.Context, userID uuid.UUID, login, firstName, lastName *string) (*repository.User, error) {
	user, err := c.repo.Users().Update(ctx, userID, login, firstName, lastName, nil)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Controller) UpdateUserPassword(ctx context.Context, userID uuid.UUID, oldPassword, newPassword string) (*repository.User, error) {
	user, err := c.repo.Users().GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	ok, err := c.hashGenerator.Compare(oldPassword, user.Password)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrIncorrectPassword
	}

	password, err := c.hashGenerator.Generate(newPassword)
	if err != nil {
		return nil, err
	}

	user, err = c.repo.Users().Update(ctx, userID, nil, nil, nil, &password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Controller) BindUserToRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID, binding Binding) error {
	var err error
	switch binding {
	case BIND:
		err = c.repo.Users().BindRole(ctx, userID, roleID)
	default:
		err = c.repo.Users().UnbindRole(ctx, userID, roleID)
	}
	return err
}

func (c *Controller) UserRoles(ctx context.Context, userID uuid.UUID) ([]*repository.Role, error) {
	return c.repo.Users().Roles(ctx, userID)
}

func (c *Controller) validateToken(ctx context.Context, token string, tokenType TokenType) (uuid.UUID, error) {

	session, err := c.repo.Sessions().GetByToken(ctx, token)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return uuid.UUID{}, ErrInvalidToken
		}
		return uuid.UUID{}, err
	}

	tokenExpired := func(ttl int) bool {
		if time.Now().Unix() > session.UpdatedAt.Unix()+int64(ttl) {
			return false
		}
		return true
	}

	ok := false
	switch tokenType {
	case RefreshToken:
		ok = tokenExpired(session.RefreshTokenTTL)
	case AccessToken:
		ok = tokenExpired(session.AccessTokenTTL)
	}

	if !ok {
		return uuid.UUID{}, ErrTokenExpired
	}

	return session.UserID, nil
}
