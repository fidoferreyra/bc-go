package main

const (
	MinTax    = 0
	MediumTax = 0.17
	MaxTax    = 0.27
)

func getTax(salary float64) (tax float64) {
	switch {
	case salary > 150000:
		tax = salary * MaxTax
	case salary > 50000:
		tax = salary * MediumTax
	default:
		tax = salary * MinTax
	}
	return tax
}
