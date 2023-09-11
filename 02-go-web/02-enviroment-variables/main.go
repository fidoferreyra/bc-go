package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar las variables de entorno.")
	}

	user := os.Getenv("MY_USER")
	pass := os.Getenv("MY_PASS")

	fmt.Printf("El usuario %s fue sacado de las variables de entorno. \n", user)
	fmt.Printf("La password %s fue sacada de las variables de entorno. \n", pass)

}
