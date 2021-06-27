package http

import (
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/usecase/checkout"
	"github.com/gin-gonic/gin"
)

type HTTPDelivery struct {
	ucCheckout checkout.Usecase
}

func New(ucCheckout checkout.Usecase) HTTPDelivery {
	return HTTPDelivery{
		ucCheckout: ucCheckout,
	}
}

func (h *HTTPDelivery) SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/checkout", h.checkoutHTTPHandler)
	return router
}
