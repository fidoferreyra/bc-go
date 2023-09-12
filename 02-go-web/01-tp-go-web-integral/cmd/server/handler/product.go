package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/domain"
	"github.com/bootcamp-go/Consignas-Go-Web.git/internal/product/interfaces"
	"github.com/bootcamp-go/Consignas-Go-Web.git/pkg/web"
	"github.com/gin-gonic/gin"
)

const (
	DateLayout = "02/01/2006"
)

type productHandler struct {
	s interfaces.IService
}

// NewProductHandler crea un nuevo controller de productos
func NewProductHandler(s interfaces.IService) *productHandler {
	return &productHandler{
		s: s,
	}
}

// GetAll obtiene todos los productos
func (h *productHandler) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, _ := h.s.GetAll()
		web.Success(c, http.StatusOK, products)
	}
}

// GetByID obtiene un producto por su id
func (h *productHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		product, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("product not found"))
			return
		}
		web.Success(c, http.StatusOK, product)
	}
}

// Search busca un producto por precio mayor a un valor
func (h *productHandler) Search() gin.HandlerFunc {
	return func(c *gin.Context) {
		priceParam := c.Query("priceGt")
		price, err := strconv.ParseFloat(priceParam, 64)
		if err != nil {
			web.Failure(c, http.StatusBadRequest, errors.New("invalid price"))
			return
		}
		products, err := h.s.SearchPriceGt(price)
		if err != nil {
			web.Failure(c, http.StatusNotFound, errors.New("no products found"))
			return
		}
		web.Success(c, http.StatusOK, products)
	}
}

// validateEmptys valida que los campos no esten vacios
func validateEmptys(product *domain.Product) (bool, error) {
	switch {
	case product.Name == "" || product.CodeValue == "" || product.Expiration == "":
		return false, errors.New("fields can't be empty")
	case product.Quantity <= 0 || product.Price <= 0:
		if product.Quantity <= 0 {
			return false, errors.New("quantity must be greater than 0")
		}
		if product.Price <= 0 {
			return false, errors.New("price must be greater than 0")
		}
	}
	return true, nil
}

// isExpirationDateValid valida que la fecha de expiracion sea valida
func isExpirationDateValid(exp string) (bool, error) {
	_, err := time.Parse(DateLayout, exp)
	if err != nil {
		fmt.Println("Fecha inválida:", err)
		return false, err
	}
	return true, nil
}

// Post crear un producto nuevo
func (h *productHandler) Post() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product domain.Product
		err := ctx.ShouldBindJSON(&product)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid product"))
			return
		}
		valid, err := validateEmptys(&product)
		if !valid {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		valid, err = isExpirationDateValid(product.Expiration)
		if !valid {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		p, err := h.s.Create(product)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusCreated, p)
	}
}

// Delete elimina un producto
func (h *productHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		web.Success(ctx, http.StatusOK, nil)
	}
}

// Put actualiza un producto
func (h *productHandler) Put() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		var product domain.Product
		err = ctx.ShouldBindJSON(&product)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid product"))
			return
		}
		valid, err := validateEmptys(&product)
		if !valid {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		valid, err = isExpirationDateValid(product.Expiration)
		if !valid {
			web.Failure(ctx, http.StatusBadRequest, err)
			return
		}
		p, err := h.s.Update(id, product)
		if err != nil {
			web.Failure(ctx, http.StatusConflict, err)
			return
		}
		web.Success(ctx, http.StatusOK, p)
	}
}

// Patch update selected fields of a product WIP
func (h *productHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name        string  `json:"name,omitempty"`
		Quantity    int     `json:"quantity,omitempty"`
		CodeValue   string  `json:"code_value,omitempty"`
		IsPublished bool    `json:"is_published,omitempty"`
		Expiration  string  `json:"expiration,omitempty"`
		Price       float64 `json:"price,omitempty"`
	}
	return func(ctx *gin.Context) {
		var r Request
		idParam := ctx.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid request"))
			return
		}
		update := domain.Product{
			Name:        r.Name,
			Quantity:    r.Quantity,
			CodeValue:   r.CodeValue,
			IsPublished: r.IsPublished,
			Expiration:  r.Expiration,
			Price:       r.Price,
		}
		if update.Expiration != "" {
			valid, err := isExpirationDateValid(update.Expiration)
			if !valid {
				web.Failure(ctx, http.StatusBadRequest, err)
				return
			}
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(ctx, http.StatusConflict, err)
			return
		}
		web.Success(ctx, http.StatusOK, p)
	}
}
