package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

// mysqlStore Errors
var (
	ErrNotFound       = errors.New("product not found")
	ErrInternalServer = errors.New("internal server error")
)

type mysqlStore struct {
	DatabaseConnection *sql.DB
}

// Create implements StoreInterface.
func (store *mysqlStore) Create(ctx context.Context, product domain.Product) error {
	_, err := store.DatabaseConnection.ExecContext(ctx,
		"INSERT INTO products (name, quantity, code_value, is_published, price, expiration) VALUES (?, ?, ?, ?, ?, ?)",
		product.Name,
		product.Quantity,
		product.CodeValue,
		product.IsPublished,
		product.Price,
		product.Expiration)

	if err != nil {
		return ErrInternalServer
	}
	return nil
}

// Delete implements StoreInterface.
func (store *mysqlStore) Delete(ctx context.Context, id int) error {
	//Delete a row in sql database based on id from the products table
	_, err := store.DatabaseConnection.ExecContext(ctx, "DELETE FROM products WHERE id = ?", id)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNotFound
		default:
			return ErrInternalServer
		}
	}
	return nil
}

// Exists implements StoreInterface.
func (store *mysqlStore) Exists(ctx context.Context, codeValue string) bool {
	// check wether a row exists in sql database or not
	row := store.DatabaseConnection.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM products WHERE code_value = ?)", codeValue)
	// get the boolean value out of row
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

// GetById implements StoreInterface.
func (store *mysqlStore) GetById(ctx context.Context, id int) (domain.Product, error) {
	product := &domain.Product{}
	//Prepare the query
	row := store.DatabaseConnection.QueryRowContext(ctx, "SELECT * FROM products WHERE id = ?", id)

	err := row.Scan(&product.Id, &product.Name, &product.Quantity, &product.CodeValue, &product.IsPublished, &product.Expiration, &product.Price)
	if err != nil {
		fmt.Println(err.Error())
		switch err {
		case sql.ErrNoRows:
			return domain.Product{}, ErrNotFound
		default:
			return domain.Product{}, ErrInternalServer
		}
	}
	return *product, nil
}

// Update implements StoreInterface.
func (store *mysqlStore) Update(ctx context.Context, product domain.Product) error {
	// Update a row in sql database
	_, err := store.DatabaseConnection.ExecContext(ctx, "UPDATE products SET  name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?", product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price, product.Id)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return ErrNotFound
		default:
			return ErrInternalServer
		}
	}
	return nil
}

func NewMysqlStore(dbConnection *sql.DB) StoreInterface {
	return &mysqlStore{
		DatabaseConnection: dbConnection,
	}
}
