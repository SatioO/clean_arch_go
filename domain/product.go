package domain

// Product ...
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p Product) getType() CartItemType {
	return ProductType
}
