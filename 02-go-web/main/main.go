package main

import (
	"encoding/json"
	"fmt"
	entities "my-first-go-api/Entities"
	repository "my-first-go-api/Repository"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

const jsonpath = "./products.json"

var repo = repository.ProductRepository{}

func main() {
	InitializeRepository()
	//Creo el routeador
	router := gin.Default()

	// GET /ping
	router.GET("/ping", Pong)

	products := router.Group("/products")

	// GET Product
	products.GET("", Products)
	products.GET("/:id", ProductsById)
	products.GET("/search", ProductsByPrice)

	router.Run() // Iniciamos el servidor y por defecto escucha el puerto 8080

	defer fmt.Println("finaliza el programa ...")
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func Products(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, repo.GetAll())
}

func ProductsById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	// Convert id to int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "id format is not valid",
		})
		return
	}
	product, err := repo.GetById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func ProductsByPrice(ctx *gin.Context) {
	priceFilter, err := strconv.ParseFloat(ctx.DefaultQuery("priceGt", "0.0"), 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid price format.",
		})
		return
	}
	result := repo.GetByPriceGreaterThan(priceFilter)
	if len(result) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("No products with price higher than %f were found.", priceFilter),
		})
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func InitializeRepository() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	data, err := readProductsFile(jsonpath)
	if err != nil {
		panic(err)
	}
	repo.InitializeFakeDb(data)
}
func readProductsFile(filepath string) ([]entities.Product, error) {
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		return []entities.Product{}, err
	}

	//Unmarshall the JSON data into a slice
	var productsList []entities.Product
	err = json.Unmarshal(fileContent, &productsList)
	if err != nil {
		return []entities.Product{}, err
	}

	return productsList, nil
}
