package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	NAME      string `json:"nombre" binding:"required"`
	LAST_NAME string `json:"apellido" binding:"required"`
}

func main() {
	//Creo el routeador
	router := gin.Default()

	// GET /ping
	router.GET("/ping", Pong)
	// POST /saludo
	router.POST("/saludo", SayHello)

	router.Run() // Iniciamos el servidor y por defecto escucha el puerto 8080
}

func SayHello(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var persona Persona
	if err := json.Unmarshal(bodyBytes, &persona); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if persona.NAME == "" || persona.LAST_NAME == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Nombre y apellido son requeridos",
		})
		return
	}

	saludo := fmt.Sprintf("Hola %s %s!", persona.NAME, persona.LAST_NAME)
	ctx.String(http.StatusOK, saludo)
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
