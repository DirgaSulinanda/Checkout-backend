package app

import (
	checkoutHandlerUC "github.com/DirgaSulinanda/Checkout-Backend/internal/app/usecase/checkout/handler"
)

func initializeUsecases() {
	checkoutUC = checkoutHandlerUC.New(promoRepo, productRepo)
}
