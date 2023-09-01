package main

import "fmt"

func main() {

	fmt.Println("Demostracion de calculo de promedios")
	fmt.Println("Notas de alumnos : [8, 6, 4 ,9, 10, 2, 3, 4, 6, 7, 5, 4, 3, 1, 9]")
	scoring := []float64{8, 6, 4, 9, 10, 2, 3, 4, 6, 7, 5, 4, 3, 1, 9}
	average := GetAverage(scoring)

	fmt.Printf("El promedio de notas de los alumnos es de %f ", average)
}
