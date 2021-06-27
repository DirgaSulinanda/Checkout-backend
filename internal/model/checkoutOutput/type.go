package checkoutOutput

type (
	CheckoutParam struct {
		CashierName string         `json:"cashier_name"`
		Products    []ProductParam `json:"products" binding:"required"`
	}

	ProductParam struct {
		SKU      string `json:"sku"`
		Quantity int    `json:"quantity"`
	}

	CheckoutOutput struct {
		Status        string    `json:"status"`
		CashierName   string    `json:"cashier_name"`
		Products      []Product `json:"products"`
		SubTotal      float64   `json:"sub_total"`
		OriginalPrice float64   `json:"original_price"`
	}

	Product struct {
		SKU           string   `json:"sku"`
		Name          string   `json:"name"`
		Quantity      int      `json:"quantity"`
		Price         float64  `json:"price"`
		TotalPrice    float64  `json:"total_price"`
		OriginalPrice float64  `json:"original_price"`
		Promos        []string `json:"promos"`
		IsOutOfStock  bool     `json:"is_out_of_stock"`
	}
)
