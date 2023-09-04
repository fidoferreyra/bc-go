package main

const (
	ShippingPrice = 2500
)

type ProductBig struct {
	Price float64
}

func (p ProductBig) GetPrice() float64 {
	return p.Price
}

func (p ProductBig) GetFinalPrice() float64 {
	return p.Price*(1+0.06) + ShippingPrice
}
