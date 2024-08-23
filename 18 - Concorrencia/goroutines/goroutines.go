package main

import (
	"fmt"
	"time"
)

func main() {
	//Concorrencia != paralelismo
	go escrever("Olá mundo")      //go routine, independente da função escrever terminar sua execução, o programa executa as proximas linhas
	escrever("Programando em Go") // se colocar go nessa chamada, o programa vai terminar sem escrever nada na tela pois apos esta linha o programa termina
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
