package lib

import "fmt"

type ErrBelowMinimumSalary struct {
	MinimumSalary int
	Salary        int
}

func (e *ErrBelowMinimumSalary) Error() string {
	return fmt.Sprintf("Error: el salario asignado: %d no alcanza el minimo imponible", e.Salary)
}

type ErrBelow10_000Salary struct {
	Salary int
}

func (e *ErrBelow10_000Salary) Error() string {
	return fmt.Sprintf("Error: el salario asignado: %d es menor a 10000", e.Salary)
}
