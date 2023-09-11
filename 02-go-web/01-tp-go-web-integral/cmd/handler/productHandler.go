package handler

import (
	"errors"
	"fmt"
	"my-first-go-api/internal/domain"
	"my-first-go-api/internal/product"
	"my-first-go-api/pkg"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const DateLayout = "xx/xx/xxxx"

type ProductHandler struct {
	service *product.ProductService
}

func NewProduct(serv *product.ProductService) ProductHandler {
	return ProductHandler{service: serv}
}

func (p *ProductHandler) Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (p *ProductHandler) GetProducts(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, p.service.GetAll())
}

func (p *ProductHandler) GetProductById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	// Convert id to int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id format is not valid",
		})
		return
	}
	product, err := p.service.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (p *ProductHandler) GetProductsByPrice(ctx *gin.Context) {
	priceFilter, err := strconv.ParseFloat(ctx.DefaultQuery("priceGt", "0.0"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid price format.",
		})
		return
	}
	result, err := p.service.GetByPriceGreaterThan(priceFilter)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("No products with price higher than %f were found.", priceFilter),
		})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (p *ProductHandler) PostProduct(ctx *gin.Context) {
	var request domain.ProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON provided",
		})
		return
	}
	newProductDTO := pkg.ProductDTO{
		Name:         request.Name,
		Quantity:     request.Quantity,
		Code_Value:   request.Code,
		Is_Published: request.Published,
		Expiration:   request.Expiration,
		Price:        request.Price,
	}
	newProduct, err := p.service.AddProduct(newProductDTO)

	// En caso de que el producto ya exista o su expiration no sea una fecha valida
	//TODO: Corregir porque entra aca siempre
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	// Creamos el response si esta todo bien
	response := domain.ProductResponse{
		Id:         newProduct.Id,
		Name:       newProduct.Name,
		Quantity:   newProduct.Quantity,
		Code:       newProduct.Code_Value,
		Published:  newProduct.Is_Published,
		Expiration: newProduct.Expiration,
		Price:      newProduct.Price,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (p *ProductHandler) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err = p.service.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "product deleted"})
}

func (p *ProductHandler) Put(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid format id"})
		return
	}
	var product domain.Product
	err = ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid product"})
		return
	}
	err = validateEmptys(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !isValidDate(product.Expiration) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid expiration date"})
		return
	}
	productModified, err := p.service.Update(id, product)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, productModified)
}
func (p *ProductHandler) Patch(ctx *gin.Context) {
	var request domain.ProductRequest
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	update := domain.Product{
		Name:         request.Name,
		Quantity:     request.Quantity,
		Code_Value:   request.Code,
		Is_Published: request.Published,
		Expiration:   request.Expiration,
		Price:        request.Price,
	}
	if !isValidDate(update.Expiration) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "expiration date is not valid"})
		return
	}

	modifiedProduct, err := p.service.Update(id, update)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, modifiedProduct)
}
func isValidDate(expiration string) bool {
	// Parse la fecha con base en el layout definido
	_, err := time.Parse(DateLayout, expiration)
	return err == nil
}

func validateEmptys(product *domain.Product) error {
	switch {
	case product.Name == "" || product.Code_Value == "" || product.Expiration == "":
		return errors.New("fields can't be empty")
	case product.Quantity <= 0 || product.Price <= 0:
		if product.Quantity <= 0 {
			return errors.New("quantity must be greater than 0")
		}
		if product.Price <= 0 {
			return errors.New("price must be greater than 0")
		}
	}
	return nil
}
