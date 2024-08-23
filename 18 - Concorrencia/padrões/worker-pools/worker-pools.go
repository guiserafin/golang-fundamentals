package main

import "fmt"

func main() {

	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}

	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

// tarefas é um canal q so recebe dados e resultados é um canal que só envia dados, usado quando ha uma fila de tarefas a ser executada
func worker(tarefas <-chan int, resultados chan<- int) {

	for numero := range tarefas {
		resultados <- fib(numero)
	}

}

func fib(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fib(posicao-2) + fib(posicao-1)
}
