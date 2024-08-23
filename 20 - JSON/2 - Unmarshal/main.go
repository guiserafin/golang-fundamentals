package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"` //se trocar nome por "-" ele ignora o campo
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	var c cachorro

	if err := json.Unmarshal([]byte(cachorroJSON), &c); err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)

	cachorro2JSON := `{"nome":"Toby","raca":"Pastor"}`

	c2 := make(map[string]string)

	if err := json.Unmarshal([]byte(cachorro2JSON), &c2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(c2)

}
