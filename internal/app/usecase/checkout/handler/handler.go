package handler

import (
	"context"

	m "github.com/DirgaSulinanda/Checkout-Backend/internal/model/checkoutOutput"
)

func (c *checkoutUC) DoCheckout(ctx context.Context, param m.CheckoutParam) (data m.CheckoutOutput, err error) {
	err = c.safeguard(ctx)
	if err != nil {
		return
	}

	status := m.StatusSuccess
	defer func() { data.Status = status }()

	promo, productDetail, err := c.prepareItems(ctx, param)
	if err != nil {
		status = m.StatusFailed
		return
	}

	// parse promo into tmpPromo
	parsedPromo, conditionChecker := parsePromo(ctx, promo)

	// parse product detail into map product detail
	parsedPDetail := parseProductDetail(ctx, productDetail)

	data, err = countTransaction(ctx, param, parsedPDetail)
	if err != nil {
		status = m.StatusFailed
		return
	}

	data, err = applyPromo(ctx, data, parsedPromo, conditionChecker)
	if err != nil {
		status = m.StatusFailed
		return
	}

	// TODO: update db to reduce quantity

	return
}
