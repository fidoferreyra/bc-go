package product

import (
	"context"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type Service interface {
	// GetByID busca un producto por su id
	GetByID(ctx context.Context, id int) (prod domain.Product, err error)
	// Create agrega un nuevo producto
	Create(ctx context.Context, p domain.Product) (prod domain.Product, err error)
	// Delete elimina un producto
	Delete(ctx context.Context, id int) error
	// Update actualiza un producto
	Update(ctx context.Context, id int, p domain.Product) (prod domain.Product, err error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(ctx context.Context, id int) (domain.Product, error) {
	p, err := s.r.GetByID(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *service) Create(ctx context.Context, p domain.Product) (domain.Product, error) {
	p, err := s.r.Create(ctx, p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *service) Update(ctx context.Context, id int, u domain.Product) (domain.Product, error) {
	p, err := s.r.GetByID(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	if u.Name != "" {
		p.Name = u.Name
	}
	if u.CodeValue != "" {
		p.CodeValue = u.CodeValue
	}
	if u.Expiration != "" {
		p.Expiration = u.Expiration
	}
	if u.Quantity > 0 {
		p.Quantity = u.Quantity
	}
	if u.Price > 0 {
		p.Price = u.Price
	}
	p, err = s.r.Update(ctx, id, p)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.r.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
