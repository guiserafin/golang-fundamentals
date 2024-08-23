package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dálmata", 3}
	fmt.Println(c)

	cachorroJson, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJson) //[]uint8

	fmt.Println(bytes.NewBuffer(cachorroJson))

	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Pastor alemão",
	}

	cachorro2JSON, err := json.Marshal(c2)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bytes.NewBuffer(cachorro2JSON))
}
