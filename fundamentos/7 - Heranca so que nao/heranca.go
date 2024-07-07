package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint16
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Joao", "Pedro", 20, 175}
	fmt.Println(p1)

	e1 := estudante{p1, "Medicina", "UFRJ"}
	fmt.Println(e1)
	fmt.Println(e1.nome, e1.faculdade)

}
