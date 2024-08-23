package main

import "fmt"

func soma(numeros ...int) int { //vai receber de 1 n ints
	// numeros é um slice
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8))

	escrever("Olá mundo", 10, 20, 30, 40, 1, 2, 3)
}
