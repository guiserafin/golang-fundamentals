// TESTE DE UNIDADE
package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	// Começa obrigatoriamente com a palavra Test
	//e o nome da funcao a ser testada é convençao

	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo inválido"},
		{"RUA DOS BOBOS", "Rua"},
		{"Estrada da Graciosa", "Estrada"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		if verify := TipoDeEndereco(cenario.enderecoInserido); verify != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				verify,
				cenario.retornoEsperado,
			)
		}
	}

}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou")
	}
}

// go test ./... - testa todos os pacotes
// go test -v - saída verbosa
// go test --cover - mostra a cobertura dos nossos testes
// go test --coverprofile x.txt - gera um txt que da um relatorio sobre a cobertura dos testes
// go tool cover --func=x.txt - le o arquivo txt e mostra no terminal a função que ta sendo testada e sua cobertura
// go tool cover --html=x.txt - le o arquivo txt e mostra em um arquivo html todas  as linhas que não estão cobertas (massa)
