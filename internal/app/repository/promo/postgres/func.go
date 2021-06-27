package postgres

import (
	"context"

	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/promo"
)

func (p *promoRepository) GetActivePromo(ctx context.Context) (data []promo.Promo, err error) {
	err = p.safeguard(ctx)
	if err != nil {
		return
	}

	rows, err := p.db.QueryContext(ctx, getQuery())
	if err != nil {
		err = ErrQuery
		return
	}

	for rows.Next() {
		var id int
		var name, description, formula string
		var enabled bool
		err = rows.Scan(&id, &name, &description, &formula, &enabled)
		if err != nil {
			continue
		}
		data = append(data, promo.Promo{
			ID:          id,
			Name:        name,
			Description: description,
			Formula:     formula,
			Enabled:     enabled,
		})
	}
	return
}

func (p *promoRepository) safeguard(ctx context.Context) error {
	// escape when ctx.Done()
	select {
	case <-ctx.Done():
		return ErrCtxTimeout
	default:
	}

	if p.db == nil {
		return ErrNilDbConnection
	}

	return nil
}

func getQuery() string {
	return "SELECT id, name, description, formula, enabled FROM promo WHERE enabled=true"
}
