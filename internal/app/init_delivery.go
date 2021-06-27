package app

import (
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/delivery/http"
)

func initializeDeliveries() {
	httpDelivery := http.New(checkoutUC)
	router = httpDelivery.SetupRouter()
}
