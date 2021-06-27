package checkout

import (
	"context"

	m "github.com/DirgaSulinanda/Checkout-Backend/internal/model/checkoutOutput"
)

type Usecase interface {
	// DoCheckout perform checkout handler
	DoCheckout(ctx context.Context, param m.CheckoutParam) (data m.CheckoutOutput, err error)
}
