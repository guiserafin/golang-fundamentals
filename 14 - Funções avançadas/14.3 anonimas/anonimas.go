package main

import "fmt"

func main() {

	retorno := func(txt string) string { //funcao anonima declarada
		return fmt.Sprintf("recebido -> %s", txt)
	}("Olá mundo") //chamda da funcao declarada

	fmt.Println(retorno)
}
