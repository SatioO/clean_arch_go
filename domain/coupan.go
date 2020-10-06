package domain

import (
	"fmt"
)

// Coupan ...
type Coupan interface {
	getDiscountAmount(float64) float64
}

// GeneralCoupan ...
type GeneralCoupan struct {
	ID       int     `json:"id"`
	Code     string  `json:"code"`
	Discount float64 `json:"discount"`
}

func (c GeneralCoupan) getType() CartItemType {
	return CoupanType
}

func (c GeneralCoupan) getDiscountAmount(price float64) float64 {
	return (price * c.Discount) / 100
}

// NextCoupan ...
type NextCoupan struct {
	ID       int     `json:"id"`
	Code     string  `json:"code"`
	Discount float64 `json:"discount"`
}

func (c NextCoupan) getType() CartItemType {
	return CoupanType
}

func (c NextCoupan) getDiscountAmount(price float64) float64 {
	return (price * c.Discount) / 100
}

// SpecialCoupan ...
type SpecialCoupan struct {
	ID       int     `json:"id"`
	Code     string  `json:"code"`
	Discount float64 `json:"discount"`
}

func (c SpecialCoupan) getType() CartItemType {
	return CoupanType
}

func (c SpecialCoupan) getDiscountAmount(price float64) float64 {
	fmt.Println("SpecialCoupan", c.Discount, price)
	return 0.0
}
