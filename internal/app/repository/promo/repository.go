package promo

import "context"

type Repository interface {
	// GetActivePromo return active promo details
	GetActivePromo(ctx context.Context) (data []Promo, err error)
}
