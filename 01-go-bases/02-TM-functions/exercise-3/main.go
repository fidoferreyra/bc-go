package main

import "fmt"

const (
	categoryA = "A"
	categoryB = "B"
	categoryC = "C"
	hourRateA = 3000
	hourRateB = 1500
	hourRateC = 1000
	bonusA    = 0.5
	bonusB    = 0.2
)

func main() {
	fmt.Println("Calculating salaries for:")
	fmt.Println("i. Category C employee, 1800 minutes of work. Result: ", getSalary("C", 1800))
	fmt.Println("ii. Category B employee, 1800 minutes of work. Result: ", getSalary("B", 1800))
	fmt.Println("iii. Category A employee, 1800 minutes of work. Result: ", getSalary("A", 1800))
	fmt.Println("iv. Category A employee, 2400 minutes of work. Result: ", getSalary("A", 2400))
	fmt.Println("vi. Category B employee, 2400 minutes of work. Result: ", getSalary("B", 2400))
	fmt.Println("vi. Category C employee, 2400 minutes of work. Result: ", getSalary("C", 2400))
}

func getSalary(category string, workMinutes int) (salary float64) {
	var workHours float64 = float64(workMinutes / 60)
	switch category {
	case categoryA:
		return getSalaryCategoryA(workHours)
	case categoryB:
		return getSalaryCategoryB(workHours)
	case categoryC:
		return getSalaryCategoryC(workHours)
	default:
		return 0
	}
}

func getSalaryCategoryA(hours float64) float64 {
	return (1 + bonusA) * (hours * hourRateA)
}

func getSalaryCategoryB(hours float64) float64 {
	return (1 + bonusB) * (hours * hourRateB)
}

func getSalaryCategoryC(hours float64) float64 {
	return hours * hourRateC
}
