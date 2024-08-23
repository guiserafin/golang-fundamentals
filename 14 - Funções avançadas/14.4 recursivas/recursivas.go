package main

import "fmt"

func fib(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fib(posicao-2) + fib(posicao-1)
}

func main() {

	fmt.Println(fib(10))
}
