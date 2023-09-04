package main

type ProductSmall struct {
	Price float64
}

func (p ProductSmall) GetPrice() float64 {
	return p.Price
}

func (p ProductSmall) GetFinalPrice() float64 {
	return p.Price
}
