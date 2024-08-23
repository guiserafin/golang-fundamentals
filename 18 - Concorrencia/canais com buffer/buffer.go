package main

import "fmt"

func main() {
	canal := make(chan string, 2) //segundo parametro = canal com capacidade 2, so bloqueia quando ultrapassa sua capacidae mexima
	canal <- "Olá mundo"
	canal <- "Olá mundo 2"
	// canal <- "Terceiro valor" //causa deadlock

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, mensagem2)
}
