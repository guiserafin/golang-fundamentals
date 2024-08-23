package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1 // Cópia

	fmt.Println(variavel1, variavel2)

	variavel1++ //Muda so o valor de variavel1

	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REFERENCIA DE MEMÓRIA
	var variavel3 int = 100
	var ponteiro *int = &variavel3
	fmt.Println(variavel3, *ponteiro) //desreferenciação

	variavel3--
	fmt.Println(variavel3, *ponteiro) //ambos serão 99

	*ponteiro = 300
	fmt.Println(variavel3, *ponteiro) //ambos serão 300

}
