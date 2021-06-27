package postgres

import (
	"context"

	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/product"
	"github.com/jmoiron/sqlx"
)

func (p *productRepository) GetProductDetails(ctx context.Context, sku []string) (data []product.Product, err error) {
	err = p.safeguard(ctx)
	if err != nil {
		return
	}

	query, args, err := getQuery(sku)
	if err != nil {
		return
	}

	query = p.db.Rebind(query)
	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		err = ErrQuery
		return
	}

	for rows.Next() {
		var sku, name string
		var price float64
		var quantity int
		err = rows.Scan(&sku, &name, &price, &quantity)
		if err != nil {
			continue
		}
		data = append(data, product.Product{
			SKU:      sku,
			Name:     name,
			Price:    price,
			Quantity: quantity,
		})
	}
	return
}

func (p *productRepository) safeguard(ctx context.Context) error {
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

func getQuery(sku []string) (q string, args []interface{}, err error) {
	q, args, err = sqlx.In("SELECT sku, name, price, quantity FROM product WHERE sku IN (?)", sku)
	if err != nil {
		return "", nil, ErrBindQuery
	}
	return q, args, nil
}
