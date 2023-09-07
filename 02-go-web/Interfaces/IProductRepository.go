package interfaces

import entities "my-first-go-api/Entities"

type IProductRepository interface {
	GetAll() []entities.Product
	GetById(id int) (entities.Product, error)
	GetByPriceGreaterThan(price float64) []entities.Product
	AddProduct(name string, quantity int, code_value string, is_published bool, expiration string, price float64)
}
