package postgres

import (
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/product"
	"github.com/jmoiron/sqlx"
)

type productRepository struct {
	db *sqlx.DB
}

// New creates a new product repository postgres implementation
func New(db *sqlx.DB) product.Repository {
	return &productRepository{db}
}
