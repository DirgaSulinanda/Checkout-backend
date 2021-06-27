package postgres

import (
	"errors"
	"fmt"
)

var (
	// ErrCtxTimeout - error type returned if context timeout
	ErrCtxTimeout = errors.New("context timeout")

	// ErrNilDbConnection err returned if db connection is nil
	ErrNilDbConnection = fmt.Errorf("Nil DB connection")

	// ErrQuery err returned if bind query db returns error
	ErrBindQuery = fmt.Errorf("Failed to do bind query")

	// ErrQuery err returned if query db returns error
	ErrQuery = fmt.Errorf("Failed to do query into postgres")
)
