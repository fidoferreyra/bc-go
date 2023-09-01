package main

import (
	"fmt"
)

func main() {

	var (
		age               = 25
		employed          = true
		salary            = 100000
		yearsOfExperience = 3
	)
	printMessageForLoan(age, employed, salary, yearsOfExperience)
}

func printMessageForLoan(age int, employed bool, salary, yearsOfExperience int) {
	switch {
	case age <= 22 || !employed || yearsOfExperience < 2:
		fmt.Println("Usted no aplica para el prestamo.")
	case salary <= 100000:
		fmt.Println("Se le otorgara el prestamo con interes.")
	default:
		fmt.Println("Se le otorgara el prestamo SIN intereses. Felicidades!")
	}
}
