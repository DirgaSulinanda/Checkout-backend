package app

import (
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/product"
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/promo"
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/usecase/checkout"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var (
	// db connections
	dbPgClient *sqlx.DB

	// repositories
	productRepo product.Repository
	promoRepo   promo.Repository

	// usecases
	checkoutUC checkout.Usecase

	// deliveries
	router *gin.Engine
)
