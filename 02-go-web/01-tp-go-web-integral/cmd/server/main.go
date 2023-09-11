package main

import (
	"fmt"
	handler "my-first-go-api/cmd/handler"
	product "my-first-go-api/internal/product"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := product.NewRepository()
	service := product.NewService(repo)
	p := handler.NewProduct(service)

	//Creo el routeador
	router := gin.Default()

	// GET /ping
	router.GET("/ping", p.Pong)

	products := router.Group("/products")
	{
		products.GET("", p.GetProducts)
		products.GET("/:id", p.GetProductById)
		products.GET("/search", p.GetProductsByPrice)
		products.POST("", p.PostProduct)
		products.DELETE(":id", p.Delete)
		products.PATCH(":id", p.Patch)
		products.PUT(":id", p.Put)
	}

	router.Run() // Iniciamos el servidor y por defecto escucha el puerto 8080

	defer fmt.Println("finaliza el programa ...")
}
