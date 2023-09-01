package main

func GetAverage(scorings []float64) (average float64) {
	sum := 0.0
	for score := range scorings {
		sum += scorings[score]
	}
	return sum / float64(len(scorings))
}
