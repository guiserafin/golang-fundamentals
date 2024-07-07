package main

import "fmt"

var n int

func init() { // init é um função que é executada antes da função main
	fmt.Println("executando a função init")
	n = 10
}

func main() {
	fmt.Println("Executando a função main")
	fmt.Println(n)
}
