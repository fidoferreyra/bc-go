package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTaxSalaryHigherThan150000(t *testing.T) {

	//Arrange
	var salary float64 = 200000
	expectedTax := salary * float64(MaxTax)

	//Act

	resultTax := getTax(salary)
	//Assert
	assert.Equal(t, expectedTax, resultTax)
}

func TestGetTaxSalaryBetween50000And150000(t *testing.T) {

	//Arrange
	var salary float64 = 100000
	expectedTax := salary * float64(MediumTax)

	//Act

	resultTax := getTax(salary)
	//Assert
	assert.Equal(t, expectedTax, resultTax)
}

func TestGetTaxSalaryLowerThan50000(t *testing.T) {

	//Arrange
	var salary float64 = 20000
	expectedTax := salary * float64(MinTax)

	//Act

	resultTax := getTax(salary)
	//Assert
	assert.Equal(t, expectedTax, resultTax)
}
