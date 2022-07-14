package ui

import (
	"errors"
	"main/models"
)

func ComissionSeller(plant models.Plant) (float32, error) {
	var comission float32

	if 0 < plant.Price {
		comission = plant.Price * 5 / 100
		return comission, nil
	}

	return 0, errors.New("Não há comissão.")

}
