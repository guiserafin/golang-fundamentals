package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

// usuarios tem o metodo salvar - u é uma variavel q podemos usar para referenciar outros campos do usuario que chamou o metodo
func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados \n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"User 1", 22}
	usuario1.salvar()

	usuario2 := usuario{"Marcelo", 30}
	usuario2.salvar()

	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
