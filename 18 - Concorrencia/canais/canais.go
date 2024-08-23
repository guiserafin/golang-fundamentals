package main

import (
	"fmt"
	"time"
)

func main() {
	//canais enviam e recebem dados

	canal := make(chan string)

	go escrever("Olá mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		// espera o canal canal receber um valor
		mensagem, aberto := <-canal //Programa é "Bloqueado" aqui até o canal receber um valor
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	} // Causa um deadlock = Nao há nenhuma goroutine enviando dados para o canal mas o canal ainda espera receber (se o canal nao tiver o fechado e nao tiver o break no laço)

	for mensagem := range canal { // faz a mesma coisa que o for acima (sintaxe simplificada)
		fmt.Println(mensagem)
	}
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal) //fecha o canal
}
