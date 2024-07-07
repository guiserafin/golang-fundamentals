package main

import "fmt"

func main() {
	fmt.Println("ESTRUTURAS DE CONTROLE")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 { //declara a variavel outroNumero ja dentro do if
		fmt.Println(" número é maior do que 0")
	} else if numero < -10 {
		fmt.Println("Numero é menor que 10")
	} else {
		fmt.Println("Numero está entre -10 e 0")
	}

	//fmt.Println(outroNumero) // ERRO - variavel OutroNumero está limitada ao escopo do if
}
