package main

import (
	"errors"
	"errors/lib"
	"fmt"
)

var errBelow10000 = lib.ErrBelow10_000Salary{}
var ErrBelowMinimumSalary = lib.ErrBelowMinimumSalary{}

func main() {
	salaryBelow := 9_000
	salary := 135_000
	salaryValid := 160_000

	_, err := CheckSalaryFoTaxes(salaryBelow)

	if errors.Is(err, &errBelow10000) {
		fmt.Println("El error arrojado es del tipo ErrBelowTenThousand")
	}
	PrintValidation(CheckSalaryFoTaxes(salary))
	PrintValidation(CheckSalaryFoTaxes(salaryValid))

}

func CheckSalaryFoTaxes(salary int) (string, error) {
	if salary < 10_000 {
		return "", &errBelow10000
	}
	if salary < 150_000 {
		return "", &lib.ErrBelowMinimumSalary{
			Salary: salary,
		}
	}
	return fmt.Sprintf("El salario de %d debe pagar impuesto", salary), nil
}
func PrintValidation(message string, e error) {
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(message)
	}
}
