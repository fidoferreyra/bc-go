package main

const (
	Tarantule     = "Tarantula"
	Cat           = "Gato"
	Hamster       = "Hamster"
	Dog           = "Perro"
	TarantuleFood = 0.15000
	hamsterFood   = 0.25000
	catFood       = 5
	dogFood       = 10
)

func operacion(animal string) (func(quantity float64) float64, string) {
	switch animal {
	case Tarantule:
		return calculateTarantule, ""
	case Cat:
		return calculateCat, ""
	case Hamster:
		return calculateHamster, ""
	case Dog:
		return calculateDog, ""
	default:
		return nil, "el animal ingresado no existe"
	}
}

func calculateTarantule(quantity float64) float64 {
	return quantity * float64(TarantuleFood)
}

func calculateCat(quantity float64) float64 {
	return quantity * float64(catFood)
}

func calculateHamster(quantity float64) float64 {
	return quantity * float64(hamsterFood)
}

func calculateDog(quantity float64) float64 {
	return quantity * float64(dogFood)
}
