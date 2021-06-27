package app

import (
	productPgRepo "github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/product/postgres"
	promoPgRepo "github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/promo/postgres"
)

func initializeRepositories() {
	productRepo = productPgRepo.New(dbPgClient)
	promoRepo = promoPgRepo.New(dbPgClient)
}
