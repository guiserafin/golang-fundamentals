package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	TipoDeEndereco := enderecos.TipoDeEndereco("Avenida Paulista")
	fmt.Println(TipoDeEndereco)
}
