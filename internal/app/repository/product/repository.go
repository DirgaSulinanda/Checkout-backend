package product

import "context"

type Repository interface {
	// GetProductDetails return product details from given sku
	GetProductDetails(ctx context.Context, sku []string) (data []Product, err error)
}
