package main

import (
	"database/sql"

	"github.com/bootcamp-go/consignas-go-db.git/cmd/server/handler"
	"github.com/bootcamp-go/consignas-go-db.git/internal/product"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func init() {}
func main() {
	dsn := mysql.Config{
		User:   "user1",
		Passwd: "secret_password",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "my_db",
	}
	database, err := sql.Open("mysql", dsn.FormatDSN())
	if err != nil {
		panic(err)
	}

	defer database.Close()

	if err = database.Ping(); err != nil {
		panic(err)
	}

	mysqlStore := store.NewMysqlStore(database)
	repo := product.NewRepository(mysqlStore)
	service := product.NewService(repo)
	productHandler := handler.NewProductHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	products := r.Group("/products")
	{
		products.GET(":id", productHandler.GetByID())
		products.POST("", productHandler.Post())
		products.DELETE(":id", productHandler.Delete())
		products.PATCH(":id", productHandler.Patch())
		products.PUT(":id", productHandler.Put())
	}

	r.Run(":8080")
}
