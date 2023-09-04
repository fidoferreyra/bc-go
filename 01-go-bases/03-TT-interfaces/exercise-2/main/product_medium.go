package main

type ProductMedium struct {
	Price float64
}

func (p ProductMedium) GetPrice() float64 {
	return p.Price
}

func (p ProductMedium) GetFinalPrice() float64 {
	return p.Price * (1 + 0.03)
}
