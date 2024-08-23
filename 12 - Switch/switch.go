package main

import "fmt"

func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado"
	default:
		return "Inválido"
	}

}

// Outra forma de fazer o switch
func diaSemana2(numero int) string {
	var diaSemana string
	switch {
	case numero == 1:
		diaSemana = "Domingo"
		fallthrough // contrário do break
	case numero == 2:
		diaSemana = "Segunda-feira"
	case numero == 3:
		diaSemana = "Terça-feira"
	default:
		diaSemana = "Inválido"
	}
	return diaSemana
}

func main() {
	fmt.Println("SWITCH")

	dia := diaSemana(3)

	fmt.Println(dia)

	dia2 := diaSemana2(2)
	fmt.Println(dia2)
}
