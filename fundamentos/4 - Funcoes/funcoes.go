package main

import (
	"fmt"
	"reflect"
)

// Funcoes também são tipos em Go

// somar(parametros) tipo_do_retorno
func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// ----------------------------------- 2 retornos do tipo int8
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(10, 20)

	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	f("Função f")
	fmt.Println(reflect.TypeOf(f)) //Imprime func()

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)

	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma2, _ := calculosMatematicos(2, 4) //_ é pra quando eu nao quero o n-ésimo retorno

	fmt.Println(resultadoSoma2)
}
