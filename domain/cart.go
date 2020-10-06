package domain

import (
	"log"
)

// CartItemType ...
type CartItemType string

const (
	// ProductType ...
	ProductType CartItemType = "Product"
	// CoupanType ...
	CoupanType CartItemType = "Coupan"
)

// CartItem ...
type CartItem interface {
	getType() CartItemType
}

// Cart ...
type Cart struct {
	Products map[string]Product   `json:"products"`
	Coupans  map[Product][]Coupan `json:"coupans"`
}

// Add puts the cart items into specific buckets(coupans, products)
func (c *Cart) Add(items []CartItem) {
	generalCoupans := c.getAllGeneralCoupans(items)
	nextCoupans := []Coupan{}

	c.Products = map[string]Product{}
	c.Coupans = map[Product][]Coupan{}

	for _, v := range items {
		switch v.(type) {
		case GeneralCoupan:
			continue

		case NextCoupan:
			nextCoupan := v.(NextCoupan)
			nextCoupans = append(nextCoupans, nextCoupan)

		case SpecialCoupan:
			continue

		case Product:
			coupans := []Coupan{}
			// append general coupans
			coupans = append(coupans, generalCoupans...)
			// append next coupans
			coupans = append(coupans, nextCoupans...)
			nextCoupans = nextCoupans[:0]
			// append products and coupans
			product := v.(Product)
			c.Products[product.Name] = product
			c.Coupans[product] = coupans

		default:
			log.Fatal("unexpected type: ", v)
		}
	}
}

// TotalPrice calculates total price based on discount
func (c *Cart) getAllGeneralCoupans(items []CartItem) []Coupan {
	coupans := []Coupan{}

	for _, v := range items {
		if v.getType() == CoupanType {
			switch v.(type) {
			case GeneralCoupan:
				coupans = append(coupans, v.(GeneralCoupan))
			}
		}
	}

	return coupans
}

// TotalPrice ...
func (c *Cart) TotalPrice() float64 {
	total := 0.0
	for _, product := range c.Products {
		productCost := product.Price
		for _, coupan := range c.Coupans[product] {
			productCost -= coupan.getDiscountAmount(productCost)
		}

		total += productCost
	}

	return total
}
