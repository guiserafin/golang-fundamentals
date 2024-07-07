package main

import "fmt"

// retorna uma função
func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()

	funcaoNova()
}
