package auxiliar

import "fmt"

// Escreve algo na tela (documentacao)
func Escrever() {
	fmt.Println("Escrevendo do pacote auxiliar")
	escrever2() //essa funcao pode ser usada pois esta no mesmo pacote (letra minuscula)
}
