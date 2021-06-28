package handler

import (
	"errors"
)

var (
	// ErrCtxTimeout - error type returned if context timeout
	ErrCtxTimeout = errors.New("context timeout")

	// ErrNilProductRepo err returned if product repo is nil
	ErrNilProductRepo = errors.New("nil product repo")

	// ErrNilPromoRepo err returned if promo repo is nil
	ErrNilPromoRepo = errors.New("nil promo repo")
)
