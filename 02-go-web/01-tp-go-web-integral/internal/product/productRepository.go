package product

import (
	"encoding/json"
	"errors"
	"fmt"
	domain "my-first-go-api/internal/domain"
	"my-first-go-api/pkg"
	"os"
)

const jsonpath = "./products.json"

type ProductRepository struct {
	productsDb []domain.Product
}

func NewRepository() *ProductRepository {
	repository := &ProductRepository{}
	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	data, err := readProductsFile(jsonpath)
	if err != nil {
		panic(err)
	}
	repository.InitializeFakeDb(data)
	return repository
}

func (repo *ProductRepository) GetAll() []domain.Product {
	return repo.productsDb
}

func (repo *ProductRepository) GetById(id int) (domain.Product, error) {
	for _, product := range repo.productsDb {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")
}
func (repo *ProductRepository) GetByCode(code string) (domain.Product, error) {
	for _, product := range repo.productsDb {
		if product.Code_Value == code {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")
}

func (repo *ProductRepository) GetByPriceGreaterThan(price float64) []domain.Product {
	var result []domain.Product
	for _, product := range repo.productsDb {
		if product.Price > price {
			result = append(result, product)
		}
	}
	return result
}

func (repo *ProductRepository) AddProduct(prod pkg.ProductDTO) domain.Product {
	var newProduct = domain.Product{
		Id:           repo.getLastId() + 1,
		Name:         prod.Name,
		Quantity:     prod.Quantity,
		Code_Value:   prod.Code_Value,
		Is_Published: prod.Is_Published,
		Expiration:   prod.Expiration,
		Price:        prod.Price,
	}
	repo.productsDb = append(repo.productsDb, newProduct)
	return newProduct
}

func (repo *ProductRepository) Delete(id int) error {
	for index, product := range repo.productsDb {
		if product.Id == id {
			repo.productsDb = append(repo.productsDb[:index], repo.productsDb...)
			return nil
		}
	}
	return errors.New("product not found")
}

func (repo *ProductRepository) Update(id int, update domain.Product) (domain.Product, error) {
	for i, product := range repo.productsDb {
		if product.Id == id {
			repo.productsDb[i] = update
			return update, nil
		}
	}
	return domain.Product{}, errors.New("product not found")
}

func (repo *ProductRepository) getLastId() int {
	return repo.productsDb[len(repo.productsDb)-1].Id
}

func (repo *ProductRepository) InitializeFakeDb(data []domain.Product) {
	repo.productsDb = data
	fmt.Println("La base de datos mock ha sido inicializada.")
}

func readProductsFile(filepath string) ([]domain.Product, error) {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		return []domain.Product{}, err
	}

	//Unmarshall the JSON data into a slice
	var productsList []domain.Product
	err = json.Unmarshal(fileContent, &productsList)
	if err != nil {
		return []domain.Product{}, err
	}

	return productsList, nil
}
