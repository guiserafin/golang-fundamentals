package main

import (
	"fmt"
)

func main() {

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		// time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j")
		// time.Sleep(time.Second)
		fmt.Println(j)
	}

	nomes := [3]string{"Joao", "Guilherme", "Liandra"}

	for index, val := range nomes {
		fmt.Println(index, val)
	}
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for _, letra := range "PALAVRA" {
		fmt.Println(string(letra)) //letra = codigo da tabela ascii da letra
	}

	usuario := map[string]string{
		"Nome":      "Guilherme",
		"Sobrenome": "Serafin",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor) // printa em ordem aleatoria qunado é map
	}

	// Não dá pra usar range em struct

	for { // loop infinino
		fmt.Println("Executando infinatamente")
	}
}
