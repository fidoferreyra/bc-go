package product

import (
	"context"
	"errors"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/store"
)

type Repository interface {
	// GetByID busca un producto por su id
	GetByID(ctx context.Context, id int) (domain.Product, error)
	// Create agrega un nuevo producto
	Create(ctx context.Context, p domain.Product) (domain.Product, error)
	// Update actualiza un producto
	Update(ctx context.Context, id int, p domain.Product) (domain.Product, error)
	// Delete elimina un producto
	Delete(ctx context.Context, id int) error
}

type repository struct {
	storage store.StoreInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage store.StoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(ctx context.Context, id int) (domain.Product, error) {
	product, err := r.storage.GetById(ctx, id)
	if err != nil {
		return domain.Product{}, errors.New("product not found")
	}
	return product, nil

}

func (r *repository) Create(ctx context.Context, p domain.Product) (domain.Product, error) {
	if !r.storage.Exists(ctx, p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	err := r.storage.Create(ctx, p)
	if err != nil {
		return domain.Product{}, errors.New("error creating product")
	}
	return p, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	err := r.storage.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(ctx context.Context, id int, p domain.Product) (domain.Product, error) {
	if !r.storage.Exists(ctx, p.CodeValue) {
		return domain.Product{}, errors.New("code value already exists")
	}
	err := r.storage.Update(ctx, p)
	if err != nil {
		return domain.Product{}, errors.New("error updating product")
	}
	return p, nil
}
