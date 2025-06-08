package entities_test

import (
	"testing"
)

func TestNewCPFValido(t *testing.T) {
	cpf := "529.982.247-25"
	result, err := Cpf.NewCPF(cpf)

	if err != nil {
		t.Errorf("Esperava CPF válido, mas retornou erro: %v", err)
	}
	if result.String() != "52998224725" {
		t.Errorf("Formato esperado sem pontuação, mas foi: %s", result.String())
	}
}

func TestNewCPFInvalido(t *testing.T) {
	cpf := "111.111.111-11"

	_, err := value.NewCPF(cpf)

	if err == nil {
		t.Errorf("Esperava erro para CPF inválido, mas não retornou erro")
	}
}
