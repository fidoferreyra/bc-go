package main

import "fmt"

var (
	lataAtun = Product{
		ID:          1,
		Name:        "Atun",
		Price:       125.50,
		Description: "Lata de lomito de atún marca COTO",
		Category:    "Enlatados",
	}
	manzana = Product{
		ID:          2,
		Name:        "Manzana Red Delicious",
		Price:       20.00,
		Description: "Manzana de color rojo intenso jugosa",
		Category:    "Frutas y Verduras",
	}
	laptop = Product{
		ID:          3,
		Name:        "Laptop",
		Price:       1000.50,
		Description: "Un portátil potente",
		Category:    "Electrónica",
	}

	taza = Product{
		ID:          4,
		Name:        "Taza de Café",
		Price:       15.00,
		Description: "Una taza para tu café",
		Category:    "Cocina",
	}

	libro = Product{
		ID:          5,
		Name:        "Libro: Fundamentos de Golang",
		Price:       20.00,
		Description: "Una introducción al lenguaje de programación Go",
		Category:    "Libros",
	}
)

var Products = []Product{lataAtun, manzana, laptop, taza}

func main() {

	fmt.Println("Productos guardados actualmente:")
	GetAll()

	fmt.Println("Agrego al producto libro:")
	libro.Save()
	GetAll()

	fmt.Println("Obtengo el producto con id 3:")
	result, err := GetById(3)
	if err != nil {
		fmt.Println(err)
	} else {
		result.Print()
	}
}
