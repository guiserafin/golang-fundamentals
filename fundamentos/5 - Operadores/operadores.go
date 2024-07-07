package main

import "fmt"

func main() {
	// Aritmeticos
	soma := 1 + 2
	subtracao := 2 - 4
	divisao := 10 / 2
	multiplicacao := 2 * 8
	resto := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, resto)

	/*
	 * const numero1 int16 = 32
	 * const numero2 int32 = 12
	 * soma = numero1 + numero2 //Erro pois sao de tipos diferentes, mesmo ambos sendo int
	 */

	//Atribuição
	var variavel1 string = "String"
	variavel2 := "String2"

	fmt.Println(variavel1, variavel2)

	// Operadores relacionais
	// 1 > 2
	// 1 >=2
	// 1 == 2
	// 1 <= 2
	// 1 < 2
	// 1 != 2

	// Operadores logicos
	// && and
	// || or
	// ! not

	//Operadores unários
	numero := 10
	numero++
	fmt.Println(numero)
	numero += 10
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	numero -= 25
	fmt.Println(numero)
	numero *= 5
	numero /= 5

	//Go não possui operadores ternários

}
