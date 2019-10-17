package controller

import "errors"

var (
	ErrBadCredentials    = errors.New("invalid login or password")
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrInvalidToken      = errors.New("invalid token")
	ErrTokenExpired      = errors.New("token expired")
)
