package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	// Testxxxxxx
	enderecoParaTeste := "Rua ABC"

	tipoDeEnderecoEsperado := "Rua"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
			tipoDeEnderecoEsperado,
			tipoDeEnderecoRecebido,
		)
	}

}
