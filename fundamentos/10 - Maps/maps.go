package main

import "fmt"

func main() {
	fmt.Println("MAPS")

	// Chaves e valores tem que ter todos os mesmos tipos
	// ----------tipo-chave tipo valores
	usuario := map[string]string{
		"nome":      "guilherme",
		"sobrenome": "serafin",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{ //map aninhado
		"nome": {
			"primeiro": "João",
			"segundo":  "Pedro",
		},
		"curso": {
			"nome":   "engenharia",
			"campus": "campus 1",
		},
	}

	fmt.Println(usuario2["curso"]["nome"])

	delete(usuario2, "curso")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "libra",
	}
	fmt.Println(usuario2)

}
