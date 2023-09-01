package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAverage(t *testing.T) {
	//Arrange
	scorings := []float64{8, 6, 4, 9, 10, 2, 3, 4, 6, 7, 5, 4, 3, 1, 9}
	expectedAverage := 5.4

	//Act
	resultAverage := GetAverage(scorings)

	//Assert
	assert.Equal(t, expectedAverage, resultAverage)

}
