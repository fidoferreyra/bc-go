package main

import (
	"errors"
	"fmt"
)

const (
	Small  = "SMALL"
	Medium = "MEDIUM"
	Big    = "BIG"
)

type Product interface {
	GetPrice() float64
	GetFinalPrice() float64
}

func Price(p Product) {
	fmt.Printf("El precio total del producto es %v pesos \n", p.GetFinalPrice())
}

func NewProduct(category string, price float64) (Product, error) {
	switch category {
	case Small:
		return newProductSmall(price), nil
	case Medium:
		return newProductMedium(price), nil
	case Big:
		return newProductBig(price), nil
	default:
		return nil, errors.New("the category selected does not exist")
	}
}

func newProductSmall(price float64) Product {
	p := ProductSmall{
		Price: price,
	}
	return p
}

func newProductMedium(price float64) Product {
	p := ProductMedium{
		Price: price,
	}
	return p
}

func newProductBig(price float64) Product {
	p := ProductBig{
		Price: price,
	}
	return p
}
