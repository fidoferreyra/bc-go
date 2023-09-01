package main

import "fmt"

func main() {

	word := "Cabezon"
	length := countCharactersInWord(word)
	fmt.Printf("La palabra %v contiene %v caracteres.", word, length)
	fmt.Println()
	fmt.Println("printing letters using len()")
	printLetters(word)
	fmt.Println("printing letters WITHOUT using len()")
	printLettersWithoutLen(word)
}

func printLetters(word string) {
	runeChars := []rune(word)
	for i := 0; i < len(runeChars); i++ {
		fmt.Println(string(runeChars[i])) // Si no lo casteo a String lo que va a imprimir sera el valor en unicode del caracter
	}
}

func printLettersWithoutLen(word string) {
	for _, char := range word {
		fmt.Println(string(char))
	}
}

func countCharactersInWord(word string) int {
	quantity := 0
	for i := 0; i < len(word); i++ {
		quantity++
	}
	return quantity
}
