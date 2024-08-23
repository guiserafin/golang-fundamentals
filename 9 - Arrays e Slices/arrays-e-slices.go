package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Arrays e slices")

	var array1 [5]string
	array1[0] = "Pos 0"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Pos2", "Pos 3", "Pos 4", "Pos 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // "..." Fixa o tamanho do array de acordo com a quantidade de valores passada
	fmt.Println(array3)

	slice1 := []int{10, 11, 12, 13, 14, 15, 16, 17} // Slice nao tem um tamanho fixo -- slice aponta para um array
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1), reflect.TypeOf(array1))

	slice1 = append(slice1, 18) //append retorna um slice novo
	fmt.Println(slice1)

	slice2 := array2[1:3] // 1(inclusivo) : 3(exclusivo) -- slice a partir de um array que ja existia (PONTEIRO)
	fmt.Println(slice2)   // Pos 2 Pos 3

	array2[1] = "Posição alterada"
	fmt.Println(slice2) //Posição alterada Pos 3

	fmt.Println("-----------------")
	// ARRAYS INTERNOS
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho = 10
	fmt.Println(cap(slice3)) // Capacidade = 11

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6) //Ultrapassou a capacidade = cria um outro array de 12 e dobra a capacidade

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Tamanho = 12
	fmt.Println(cap(slice3)) // Capacidade = 24

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // Tamanho = 5
	fmt.Println(cap(slice4)) // Capacidade = 5

}
