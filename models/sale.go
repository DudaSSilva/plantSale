package models

import (
	"fmt"
)

type Sale struct {
	Name       string
	SellerCode int
	Cpf        string
}

func (s Sale) ComissionSeller(plant *Plant) float64 {
	return (5 / 100) * plant.Price
}

func (s Sale) ViewComissionSeller(plant *Plant) {
	fmt.Println("\nComissão do vendedor: ", s.ComissionSeller(plant))
}
