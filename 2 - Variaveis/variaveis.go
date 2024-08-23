package main

import "fmt"

func main() {

	var nome string = "Guilherme" //declaracao explicita

	idade := 22 //inferencia de tipo

	fmt.Println(nome, idade)

	var (
		variavel3 string = "bla"
		variavel4 int    = 2
	)

	fmt.Println(variavel3, variavel4)

	var5, var6 := "teste", 2.4

	fmt.Println(var5, var6)

	const constante1 = "String constante"

	// constante1 = 'a' erro

	//Inversão de valores
	var5, nome = nome, var5

}
