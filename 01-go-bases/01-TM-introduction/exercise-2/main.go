package main

import "fmt"

func main() {
	var (
		temperature = 23.5
		humidity    = 75
		pressure    = 1000
	)
	message := fmt.Sprintf("En Buenos Aires tenemos una temperatura de %v, con una humedad del %v porciento y una presion atmosferica de %v hectopascales", temperature, humidity, pressure)

	fmt.Println(message)
}
