package handler

import (
	"errors"
	"fmt"
)

var (
	// ErrCtxTimeout - error type returned if context timeout
	ErrCtxTimeout = errors.New("context timeout")

	// ErrNilProductRepo err returned if product repo is nil
	ErrNilProductRepo = fmt.Errorf("Nil product repo")

	// ErrNilPromoRepo err returned if promo repo is nil
	ErrNilPromoRepo = fmt.Errorf("Nil promo repo")
)
