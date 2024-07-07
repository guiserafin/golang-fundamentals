package main

import "fmt"

func funcao1() {
	fmt.Println("executando a função 1")
}

func funcao2() {
	fmt.Println("executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada, resultado será retornado!")
	defer fmt.Printf("chamadas do defer são empilhadas, essa mensagem aparecerá antes que a de cima\n")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := n1 + n2/2

	if media >= 6 {
		return true
	}

	return false

}

func main() {
	// defer funcao1() //defer = adiar -> executa no final da funcao main
	// funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))

}
