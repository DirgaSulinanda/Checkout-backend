package handler

type (
	tmpPromo struct {
		id         int
		name       string
		rawFormula string
		condition  map[string]int    // map[sku]minQty
		discount   map[string]string // map[sku]discQty
		isActive   bool
	}

	activePromoStruct struct {
		conditionsMet       int
		conditionsShouldMet int
	}
)
