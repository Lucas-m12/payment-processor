package entities

import (
	"errors"

	"github.com/google/uuid"
)

type Payment struct {
	ID          string
	Value       float64
	Barcode     string
	CustomerCPF Cpf
}

func NewPayment(value float64, barcode string, customerCPF Cpf) (*Payment, error) {
	err := validatePayment(value, barcode)
	if err != nil {
		return nil, err
	}
	return &Payment{
		ID:          uuid.New().String(),
		Value:       value,
		Barcode:     barcode,
		CustomerCPF: customerCPF,
	}, nil
}

func validatePayment(value float64, barcode string) error {
	if value <= 0 {
		return errors.New("value could not be equal to 0 or less than it")
	}
	if barcode == "" {
		return errors.New("barcode could not be empty")
	}
	return nil
}
