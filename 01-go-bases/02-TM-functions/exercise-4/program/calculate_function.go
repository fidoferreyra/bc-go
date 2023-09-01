package program

import (
	"fmt"
	"math"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func operation(operator string) (func(values ...float64) float64, error) {
	switch operator {
	case minimum:
		return getMinimum, nil
	case average:
		return getAverage, nil
	case maximum:
		return getMaximum, nil
	default:
		return nil, fmt.Errorf("the operation selected does not exist")
	}
}

func getMaximum(scoring ...float64) (maxScore float64) {
	for _, score := range scoring {
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func getAverage(scoring ...float64) (average float64) {
	sum := 0.0
	for _, score := range scoring {
		sum += score
	}
	return sum / float64(len(scoring))
}

func getMinimum(scoring ...float64) (minimum float64) {
	minimum = math.MaxInt
	for _, score := range scoring {
		if score < minimum {
			minimum = score
		}
	}
	return minimum
}
