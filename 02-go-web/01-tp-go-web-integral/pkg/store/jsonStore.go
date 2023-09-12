package store

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
)

type Store interface {
	GetAll() ([]domain.Product, error)
	GetOne(id int) (domain.Product, error)
	AddOne(product domain.Product) error
	UpdateOne(product domain.Product) error
	DeleteOne(id int) error
	saveProducts(products []domain.Product) error
	loadProducts() ([]domain.Product, error)
}

type jsonStore struct {
	pathToFile string
}

// AddOne implements Store.
func (store *jsonStore) AddOne(product domain.Product) error {
	products, err := store.loadProducts()
	if err != nil {
		return err
	}
	product.Id = len(products) + 1
	products = append(products, product)
	return store.saveProducts(products)
}

// DeleteOne implements Store.
func (store *jsonStore) DeleteOne(id int) error {
	products, err := store.loadProducts()
	if err != nil {
		return err
	}
	for index, prod := range products {
		if prod.Id == id {
			products = append(products[:index], products[index+1:]...)
			return store.saveProducts(products)
		}
	}
	return errors.New("product not found")
}

// GetAll implements Store.
func (store *jsonStore) GetAll() ([]domain.Product, error) {
	products, err := store.loadProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

// GetOne implements Store.
func (store *jsonStore) GetOne(id int) (domain.Product, error) {
	products, err := store.loadProducts()
	if err != nil {
		return domain.Product{}, err
	}
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("product not found")
}

// UpdateOne implements Store.
func (store *jsonStore) UpdateOne(product domain.Product) error {
	products, err := store.loadProducts()
	if err != nil {
		return err
	}
	for index, prod := range products {
		if prod.Id == product.Id {
			products[index] = product
			return store.saveProducts(products)
		}
	}
	return errors.New("product not found")
}

// loadProducts carga los productos desde el archivo json
func (store *jsonStore) loadProducts() ([]domain.Product, error) {
	var products []domain.Product
	file, err := os.ReadFile(store.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// saveProducts guarda los productos en un archivo json
func (store *jsonStore) saveProducts(products []domain.Product) error {
	bytes, err := json.Marshal(products)
	if err != nil {
		return err
	}
	return os.WriteFile(store.pathToFile, bytes, 0644)
}

func NewStore(path string) Store {
	return &jsonStore{
		pathToFile: path,
	}
}
