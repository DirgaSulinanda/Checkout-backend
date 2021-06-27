package postgres

import (
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/promo"
	"github.com/jmoiron/sqlx"
)

type promoRepository struct {
	db *sqlx.DB
}

// New creates a new promo repository postgres implementation
func New(db *sqlx.DB) promo.Repository {
	return &promoRepository{db}
}
