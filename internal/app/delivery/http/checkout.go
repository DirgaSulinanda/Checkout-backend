package http

import (
	"context"
	"net/http"
	"time"

	m "github.com/DirgaSulinanda/Checkout-Backend/internal/model/checkoutOutput"
	"github.com/gin-gonic/gin"
)

func (h *HTTPDelivery) checkoutHTTPHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()
	var param m.CheckoutParam
	if err := c.BindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error parse json body",
		})
		return
	}

	data, err := h.ucCheckout.DoCheckout(ctx, param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
