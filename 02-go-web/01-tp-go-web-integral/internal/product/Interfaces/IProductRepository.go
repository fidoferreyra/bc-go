package interfaces

import (
	domain "my-first-go-api/internal/domain"
	"my-first-go-api/pkg"
)

type IProductRepository interface {
	GetAll() []domain.Product
	GetById(id int) domain.Product
	GetByPriceGreaterThan(price float64) []domain.Product
	AddProduct(newProduct pkg.ProductDTO) domain.Product
}
