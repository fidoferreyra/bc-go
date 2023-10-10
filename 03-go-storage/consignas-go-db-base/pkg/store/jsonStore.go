package store

import (
	"context"
	"encoding/json"
	"errors"
	"os"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type jsonStore struct {
	pathToFile string
}

// loadProducts carga los productos desde un archivo json
func (s *jsonStore) loadProducts() ([]domain.Product, error) {
	var products []domain.Product
	file, err := os.ReadFile(s.pathToFile)
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
func (s *jsonStore) saveProducts(products []domain.Product) error {
	bytes, err := json.Marshal(products)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

// GetById busca un producto por id
func (s *jsonStore) GetById(ctx context.Context, id int) (domain.Product, error) {
	products, err := s.loadProducts()
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

// Create crea un nuevo producto
func (s *jsonStore) Create(ctx context.Context, product domain.Product) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	product.Id = len(products) + 1
	products = append(products, product)
	return s.saveProducts(products)
}

// Update actualiza un producto existente
func (s *jsonStore) Update(ctx context.Context, product domain.Product) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.Id == product.Id {
			products[i] = product
			return s.saveProducts(products)
		}
	}
	return errors.New("product not found")
}

// Delete elimina un producto por id
func (s *jsonStore) Delete(ctx context.Context, id int) error {
	products, err := s.loadProducts()
	if err != nil {
		return err
	}
	for i, p := range products {
		if p.Id == id {
			products = append(products[:i], products[i+1:]...)
			return s.saveProducts(products)
		}
	}
	return errors.New("product not found")
}

// Exists verifica si un producto existe por su código
func (s *jsonStore) Exists(ctx context.Context, codeValue string) bool {
	products, err := s.loadProducts()
	if err != nil {
		return false
	}
	for _, p := range products {
		if p.CodeValue == codeValue {
			return true
		}
	}
	return false
}

// NewJsonStore crea un nuevo store de productos
func NewJsonStore(path string) StoreInterface {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonStore{
		pathToFile: path,
	}
}
