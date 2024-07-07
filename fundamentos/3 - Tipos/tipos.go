package main

import (
	"errors"
	"fmt"
)

func main() {
	// Numeros - inteiros e reais
	// Inteiros = int - usa a arq do pc como base - int8, int16, int32 e int64 -- quantidade de bits
	// uint = unsigned int
	var numero int16 = 100
	fmt.Println(numero)

	// alias
	//rune = int32
	var numero3 rune = 123456 //rune = exatamente igual ao int32
	fmt.Println(numero3)

	//uint8 = byte
	var numero4 byte = 123
	fmt.Println(numero4)

	// Floats - float32 e float64

	var numeroReal1 float32 = 123.122

	var numeroReal2 float64 = 1233000000.34

	numeroReal3 := 3.14
	fmt.Println(numeroReal1, numeroReal2, numeroReal3)

	// STRINGS -- não tem char, usar rune

	const str string = "texto"
	// var char rune = 'a'
	char := 'a'
	fmt.Println(char) // converte para int

	var texto string
	fmt.Println(texto) //string em branco

	//BOOLEANO - bool

	var booleano bool
	fmt.Println(booleano) //padrao = false

	//ERROR - error, erro é um tipo em Go

	var erro error
	fmt.Println(erro) //padrão = nil

	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}
