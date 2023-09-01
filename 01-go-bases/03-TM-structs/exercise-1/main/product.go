package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) Print() {
	pJSON, err := json.Marshal(p)
	fmt.Println(string(pJSON))
	if err != nil {
		fmt.Println(err)
	}
}

func GetAll() {
	for _, product := range Products {
		product.Print()
	}
}

func GetById(id int) (Product, error) {
	result := Product{}
	for _, product := range Products {
		if product.ID == id {
			result = product
			return result, nil
		}
	}
	return result, errors.New("el id ingresado no corresponde a ningun producto")
}
