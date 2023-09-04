package main

import "fmt"

func main() {
	var (
		cigar, errCigar   = NewProduct(Small, 480.55)
		laptop, errLaptop = NewProduct(Medium, 559999.99)
		sofa, errSofa     = NewProduct(Big, 759999.99)
	)

	fmt.Println("Testeando los precios...")

	if errCigar != nil {
		fmt.Println(errCigar)
	} else {
		fmt.Printf("El precio original de los cigarrillos es %v \n", cigar.GetPrice())
		Price(cigar)
	}

	if errLaptop != nil {
		fmt.Println(errLaptop)
	} else {
		fmt.Printf("El precio original de la laptop es %v \n", laptop.GetPrice())
		Price(laptop)
	}

	if errSofa != nil {
		fmt.Println(errSofa)
	} else {
		fmt.Printf("El precio original del sofa es %v \n", sofa.GetPrice())
		Price(sofa)
	}
}
