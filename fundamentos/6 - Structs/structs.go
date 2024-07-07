package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint32
}

type preco struct {
	reais    int
	centavos int
}

func main() {
	usuario1 := usuario{nome: "Guilherme", idade: 22} //padrão = { 0} - valor 0 de cada um dos campos, se fossem 2 ints seria {0 0}
	fmt.Println(usuario1)

	var preco1 preco
	fmt.Println(preco1)

	preco1.reais = 10
	preco1.centavos = 99
	fmt.Println(preco1)

	usuario3 := usuario{nome: "Pedro"} //usuario que nao sei a idade
	fmt.Println(usuario3)

	enderecoEx := endereco{"Rua dos bobos", 0}
	usuario4 := usuario{"Luiz", 26, enderecoEx}
	fmt.Println(usuario4)
}
