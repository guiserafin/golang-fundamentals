package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Stirng")
	generica(3)
	generica(true)

	mapa := map[interface{}]interface{}{
		1:            "string",
		float32(100): true,
		"string":     "string",
		true:         float64(32),
	}

	fmt.Println(mapa)
}
