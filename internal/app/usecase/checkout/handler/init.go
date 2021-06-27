package handler

import (
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/product"
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/promo"
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/usecase/checkout"
)

type checkoutUC struct {
	promoRepo   promo.Repository
	productRepo product.Repository
}

// New creates a new product repository postgres implementation
func New(promoRepo promo.Repository, productRepo product.Repository) checkout.Usecase {
	return &checkoutUC{promoRepo, productRepo}
}
