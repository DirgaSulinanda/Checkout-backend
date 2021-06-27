package handler

import (
	"context"
	"strconv"
	"strings"

	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/product"
	"github.com/DirgaSulinanda/Checkout-Backend/internal/app/repository/promo"
	m "github.com/DirgaSulinanda/Checkout-Backend/internal/model/checkoutOutput"
)

func (c *checkoutUC) prepareItems(ctx context.Context, param m.CheckoutParam) (promo []promo.Promo, productDetail []product.Product, err error) {
	// get promo
	promo, err = c.promoRepo.GetActivePromo(ctx)
	if err != nil {
		return
	}

	// get product detail
	var sku = []string{}
	for _, p := range param.Products {
		sku = append(sku, p.SKU)
	}

	productDetail, err = c.productRepo.GetProductDetails(ctx, sku)
	if err != nil {
		return
	}

	return
}

func (c *checkoutUC) safeguard(ctx context.Context) error {
	// escape when ctx.Done()
	select {
	case <-ctx.Done():
		return ErrCtxTimeout
	default:
	}

	if c.promoRepo == nil {
		return ErrNilPromoRepo
	}
	if c.productRepo == nil {
		return ErrNilProductRepo
	}

	return nil
}

func parsePromo(ctx context.Context, promo []promo.Promo) (result map[int]tmpPromo, conditionChecker map[string][]int) {
	result = make(map[int]tmpPromo)
	conditionChecker = make(map[string][]int)

	for _, v := range promo {
		p := tmpPromo{
			id:         v.ID,
			name:       v.Name,
			rawFormula: v.Formula,
		}

		splitValue := strings.Split(v.Formula, "=")
		if len(splitValue) != 2 {
			continue
		}
		// left argument should be treated as condition
		p.condition = parseConditionArgument(splitValue[0], v.ID, conditionChecker)

		// right argument should be treated as discout
		p.discount = parseDiscountArgument(splitValue[1])

		result[v.ID] = p
	}
	return
}

func parseConditionArgument(formula string, id int, conditionChecker map[string][]int) (result map[string]int) {
	result = make(map[string]int)
	conditionArr := strings.Split(formula, "+")
	for _, v := range conditionArr {
		// loop for each condition
		c := strings.Split(v, "*")
		if len(c) != 2 {
			continue
		}
		qty, errParse := strconv.Atoi(c[0])
		if errParse != nil {
			continue
		}
		sku := trimBrackets(c[1])
		result[sku] = qty
		// append to checker
		conditionChecker[sku] = append(conditionChecker[sku], id)
	}
	return
}

func parseDiscountArgument(formula string) (result map[string]string) {
	result = make(map[string]string)
	conditionArr := strings.Split(formula, "+")
	for _, v := range conditionArr {
		// loop for each condition
		d := strings.Split(v, "*")
		if len(d) != 2 {
			continue
		}
		result[trimBrackets(d[1])] = d[0]
	}
	return
}

func trimBrackets(s string) string {
	s = strings.ReplaceAll(s, "{", "")
	s = strings.ReplaceAll(s, "}", "")
	return s
}

func parseProductDetail(ctx context.Context, p []product.Product) (result map[string]product.Product) {
	result = make(map[string]product.Product)
	for _, v := range p {
		result[v.SKU] = v
	}
	return
}

func countTransaction(ctx context.Context, param m.CheckoutParam, productDetail map[string]product.Product) (data m.CheckoutOutput, err error) {
	var subTotal float64

	for _, v := range param.Products {
		pDetail, exists := productDetail[v.SKU]
		if !exists {
			continue
		}

		isOutOfStock := false
		qty := v.Quantity
		if pDetail.Quantity < qty {
			qty = pDetail.Quantity
			isOutOfStock = true
		}

		totalPrice := pDetail.Price * float64(qty)
		subTotal += totalPrice

		data.Products = append(data.Products, m.Product{
			SKU:           v.SKU,
			Name:          pDetail.Name,
			Quantity:      qty,
			Price:         pDetail.Price,
			OriginalPrice: totalPrice,
			TotalPrice:    totalPrice,
			IsOutOfStock:  isOutOfStock,
		})
	}

	data.OriginalPrice = subTotal
	data.SubTotal = subTotal
	data.CashierName = param.CashierName
	return
}

func applyPromo(ctx context.Context, data m.CheckoutOutput, promo map[int]tmpPromo, conditionChecker map[string][]int) (resultData m.CheckoutOutput, err error) {
	activePromo := make(map[int]activePromoStruct)
	var totalDiscount float64

	// check every products for promo
	for _, product := range data.Products {
		// check for eligible promo
		if promoIDs, eligible := conditionChecker[product.SKU]; eligible {
			checkEveryEligiblePromo(promoIDs, promo, product, activePromo)
		}
	}

	// apply discount to every active promo
	for promoID, actPromo := range activePromo {
		// only apply discount for conditions met products
		if actPromo.conditionsMet >= actPromo.conditionsShouldMet {
			currentPromo, found := promo[promoID]
			if !found {
				continue
			}

			// check every product data for discounted product
			for i, product := range data.Products {
				discountQty, found := currentPromo.discount[product.SKU]
				if !found {
					continue
				}

				qty, multiplier, errParse := parseDiscountQtyAndMultiplier(discountQty, product.Quantity)
				if errParse != nil {
					continue
				}

				discount := countDiscount(product, qty, multiplier)
				totalDiscount += discount

				product.Promos = append(product.Promos, currentPromo.name)
				product.TotalPrice = product.OriginalPrice - discount
				data.Products[i] = product
			}
		}
	}

	data.SubTotal = data.SubTotal - totalDiscount
	return data, nil
}

func checkEveryEligiblePromo(promoIDs []int, promo map[int]tmpPromo, product m.Product, activePromo map[int]activePromoStruct) {
	for _, promoID := range promoIDs {
		// check for promo id
		if prm, found := promo[promoID]; found {
			// get minQty
			minQty := prm.condition[product.SKU]
			checkPassPromoCondition(prm, product.Quantity, minQty, activePromo)
		}
	}
}

func checkPassPromoCondition(promo tmpPromo, transactionQty, promoMinQty int, activePromo map[int]activePromoStruct) {
	// check on min qty condition
	if transactionQty >= promoMinQty {
		// check if promo already exists
		if actPromo, exists := activePromo[promo.id]; exists {
			actPromo.conditionsShouldMet++
		} else {
			activePromo[promo.id] = activePromoStruct{
				conditionsMet:       1,
				conditionsShouldMet: len(promo.condition),
			}
		}
	}
}

func parseDiscountQtyAndMultiplier(discountQty string, productQuantity int) (qty float64, multiplier int, err error) {
	// apply discount for product discount by quantity
	if strings.Contains(discountQty, "n") {
		// remove n
		qtyString := strings.ReplaceAll(discountQty, "n", "")
		qty, err = strconv.ParseFloat(qtyString, 64)
		if err != nil {
			return
		}
		multiplier = productQuantity
	} else {
		// regular discount
		qty, err = strconv.ParseFloat(discountQty, 64)
		if err != nil {
			return
		}
		multiplier = 1
	}
	return
}

func countDiscount(product m.Product, qty float64, multiplier int) float64 {
	if float64(product.Quantity) < qty {
		qty = float64(product.Quantity)
	}

	// count for discount
	return qty * float64(multiplier) * product.Price
}
