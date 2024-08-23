package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Golang")
	auxiliar.Escrever()
	//funcao com letra maiuscula = "publica", se fort letra minuscula so é visivel no pacote que ela está

	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)
}

//go build compila o codigo
