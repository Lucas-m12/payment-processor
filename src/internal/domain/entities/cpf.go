package entities

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Cpf struct {
	Value string
}

func NewCpf(value string) (*Cpf, error) {
	cleanedCpf := strings.ReplaceAll(strings.ReplaceAll(value, ".", ""), "-", "")
	isValid := isValidCpf(cleanedCpf)
	if !isValid {
		return nil, errors.New("Invalid CPF")
	}
	return &Cpf{
		Value: cleanedCpf,
	}, nil
}

func isValidCpf(cpf string) bool {
	if len(cpf) != 11 || regexp.MustCompile(`^(\d)\1+$`).MatchString(cpf) {
		return false
	}

	soma := 0
	for i := 0; i < 9; i++ {
		d, _ := strconv.Atoi(string(cpf[i]))
		soma += d * (10 - i)
	}
	d1 := (soma * 10 % 11) % 10
	if d1 != int(cpf[9]-'0') {
		return false
	}

	soma = 0
	for i := 0; i < 10; i++ {
		d, _ := strconv.Atoi(string(cpf[i]))
		soma += d * (11 - i)
	}
	d2 := (soma * 10 % 11) % 10
	return d2 == int(cpf[10]-'0')
}
