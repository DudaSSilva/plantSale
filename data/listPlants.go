package data

import (
	"errors"
	"main/models"
)

var plants []models.Plant

func SearchPlantByName(searchedName string) (*models.Plant, error) {

	for _, p := range plants {
		if p.Name == searchedName {
			return &p, nil
		}
	}

	return nil, errors.New("Planta não encontrada.")
}

func SavePlant(plant models.Plant) []models.Plant {
	plants = append(plants, plant)
	return plants
}

func ViewPlant(searchedName string) (*models.Plant, error) {

	for _, p := range plants {
		if p.Name == searchedName {
			p.ViewPlantInformation()
			return &p, nil
		}
	}

	return nil, errors.New("\nNenhuma informação para planta cadastrada.")
}
