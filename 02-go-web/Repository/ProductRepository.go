package repository

import (
	"errors"
	"fmt"
	entities "my-first-go-api/Entities"
	"time"
)

type ProductRepository struct {
	productsDb []entities.Product
}

const DateLayout = "xx/xx/xxxx"

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

func (repo *ProductRepository) AddProduct(name string, quantity int, code_value string, is_published bool, expiration string, price float64) (entities.Product, error) {

	if err := checkExpirationDate(expiration); err != nil {
		return entities.Product{}, err
	}
	if err := repo.checkUniqueCode(code_value); err != nil {
		return entities.Product{}, err
	}

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
	return newProduct, nil
}

func checkExpirationDate(expiration string) error {
	// Parse la fecha con base en el layout definido
	_, err := time.Parse(DateLayout, expiration)
	if err != nil {
		return errors.New("expiration date is not valid")
	}
	return nil
}

func (repo *ProductRepository) checkUniqueCode(code string) error {
	for _, product := range repo.productsDb {
		if product.Code_Value == code {
			return errors.New("the product code already exists")
		}
	}
	return nil
}

func (repo *ProductRepository) getLastId() int {
	return repo.productsDb[len(repo.productsDb)-1].Id
}
func (repo *ProductRepository) InitializeFakeDb(data []entities.Product) {
	repo.productsDb = data
	fmt.Println("La base de datos mock ha sido inicializada.")
}
