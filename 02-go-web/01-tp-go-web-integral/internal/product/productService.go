package product

import (
	"errors"
	domain "my-first-go-api/internal/domain"
	"my-first-go-api/pkg"
)

type ProductService struct {
	repository *ProductRepository
}

func NewService(repository *ProductRepository) *ProductService {
	return &ProductService{repository: repository}
}

func (service *ProductService) GetAll() []domain.Product {
	return service.repository.GetAll()
}

func (service *ProductService) GetById(id int) (domain.Product, error) {
	product, err := service.repository.GetById(id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (service *ProductService) GetByPriceGreaterThan(price float64) ([]domain.Product, error) {
	result := service.repository.GetByPriceGreaterThan(price)
	if len(result) == 0 {
		return result, errors.New("no se han encontrado resultados")
	}
	return result, nil
}

func (service *ProductService) AddProduct(newProduct pkg.ProductDTO) (domain.Product, error) {
	if !service.isUniqueByCode(newProduct.Code_Value) {
		return domain.Product{}, errors.New("a product with this code already exists")
	}
	result := service.repository.AddProduct(newProduct)

	return result, nil
}

func (service *ProductService) Delete(id int) error {
	err := service.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (service *ProductService) Update(id int, update domain.Product) (domain.Product, error) {
	product, err := service.repository.GetById(id)
	if err != nil {
		return domain.Product{}, err
	}
	if update.Name != "" {
		product.Name = update.Name
	}
	if update.Code_Value != "" {
		product.Code_Value = update.Code_Value
	}
	if update.Expiration != "" {
		product.Expiration = update.Expiration
	}
	if update.Quantity > 0 {
		product.Quantity = update.Quantity
	}
	if update.Price > 0 {
		product.Price = update.Price
	}
	if !service.isUniqueByCode(update.Code_Value) {
		return domain.Product{}, errors.New("code must be unique")
	}
	product, err = service.repository.Update(id, product)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (service *ProductService) isUniqueByCode(code string) bool {
	_, err := service.repository.GetByCode(code)
	return err == nil
}
