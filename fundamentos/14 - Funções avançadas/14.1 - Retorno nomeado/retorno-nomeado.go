package main

import "fmt"

// ---------------------------------- retorna uma variavel chamada soma e uma chamada subtracao
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {
	soma1, subtracao1 := calculosMatematicos(3, 5)

	fmt.Println(soma1, subtracao1)
}
