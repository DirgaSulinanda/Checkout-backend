package postgres

import (
	"errors"
)

var (
	// ErrCtxTimeout - error type returned if context timeout
	ErrCtxTimeout = errors.New("context timeout")

	// ErrNilDbConnection err returned if db connection is nil
	ErrNilDbConnection = errors.New("nil DB connection")

	// ErrQuery err returned if bind query db returns error
	ErrBindQuery = errors.New("failed to do bind query")

	// ErrQuery err returned if query db returns error
	ErrQuery = errors.New("failed to do query into postgres")
)
