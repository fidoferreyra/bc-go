package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Customer struct {
	Id          int
	Name        string
	DNI         string
	PhoneNumber string
	Address     string
}

func main() {
	fmt.Println("Iniciando el programa ...")
	filepath := "./main/customers.txt"

	ReadFile(filepath)

	defer fmt.Println("Ejecucion finalizada")
}

func ReadFile(filepath string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	data, err := os.ReadFile(filepath)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s \n", data)
}

func OpenFile(filepath string) os.File {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	file, err := os.Open(filepath)
	if err != nil {
		panic(errors.New("el archivo no fue encontrado o esta corrupto"))
	}
	defer file.Close()
	return *file
}

func extractLinesFromFile(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
