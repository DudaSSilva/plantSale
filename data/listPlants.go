package data

import (
	"errors"
	"main/models"
)

var plants []models.Plant

func SearchPlantByName(searchedCodeplant int) (*models.Plant, error) {

	for _, p := range plants {
		if p.PlantCode == searchedCodeplant {
			return &p, nil
		}
	}

	return nil, errors.New("Planta não encontrada.")
}

func SavePlant(plant models.Plant) []models.Plant {
	plants = append(plants, plant)
	return plants
}
