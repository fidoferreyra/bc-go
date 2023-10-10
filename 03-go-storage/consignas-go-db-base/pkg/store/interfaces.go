package store

import (
	"context"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type StoreInterface interface {
	// GetById devuelve un producto por su id
	GetById(ctx context.Context, id int) (domain.Product, error)
	// Create agrega un nuevo producto
	Create(ctx context.Context, product domain.Product) error
	// Update actualiza un producto
	Update(ctx context.Context, product domain.Product) error
	// Delete elimina un producto
	Delete(ctx context.Context, id int) error
	// Exists verifica si un producto existe
	Exists(ctx context.Context, codeValue string) bool
}
