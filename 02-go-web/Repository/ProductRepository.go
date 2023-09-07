package repository

import (
	"errors"
	"fmt"
	entities "my-first-go-api/Entities"
)

type ProductRepository struct {
	productsDb []entities.Product
}

func (repo *ProductRepository) GetAll() []entities.Product {
	return repo.productsDb
}

func (repo *ProductRepository) GetById(id int) (entities.Product, error) {
	for _, product := range repo.productsDb {
		if product.Id == id {
			return product, nil
		}
	}
	return entities.Product{}, errors.New("el id no corresponde a un producto")
}

func (repo *ProductRepository) GetByPriceGreaterThan(price float64) []entities.Product {
	var result []entities.Product
	for _, product := range repo.productsDb {
		if product.Price > price {
			result = append(result, product)
		}
	}
	return result
}

func (repo *ProductRepository) AddProduct(name string, quantity int, code_value string, is_published bool, expiration string, price float64) {
	var newProduct = entities.Product{
		Id:           repo.getLastId() + 1,
		Name:         name,
		Quantity:     quantity,
		Code_Value:   code_value,
		Is_Published: is_published,
		Expiration:   expiration,
		Price:        price,
	}
	repo.productsDb = append(repo.productsDb, newProduct)
}

func (repo *ProductRepository) getLastId() int {
	return repo.productsDb[len(repo.productsDb)-1].Id
}
func (repo *ProductRepository) InitializeFakeDb(data []entities.Product) {
	repo.productsDb = data
	fmt.Println("La base de datos mock ha sido inicializada.")
}
