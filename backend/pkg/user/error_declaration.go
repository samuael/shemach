package user

import "errors"

var (
	ErrInternalDatabaseError      = errors.New("internal database error")
	ErrMissingAuthorizationHeader = errors.New("missing authorization header")
	ErrInvalidSession             = errors.New("invalid session")
)
