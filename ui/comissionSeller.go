package ui

import (
	"errors"
	"fmt"
	"main/models"
)

func ComissionSeller(plant models.Plant) (float32, error) {
	var comission float32

	if plant.Price > 0 {
		comission = 5 / 100 * plant.Price
		return comission, nil
		fmt.Printf("\n %f", comission)
	}

	return 0, errors.New("Não há comissão para este vendedor.")

}
