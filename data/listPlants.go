package data

import (
  "fmt"
	"main/models"
	"errors"
)

var plants []models.Plant

func SearchPlantByName(searchedName string) (*models.Plant, error) {

	for _, p := range plants {
		// p ===== plantas[i]
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

func ViewPlant(searchedName string){

	for _, p := range plants {
		if p.Name == searchedName {
			p.ViewPlantInformation()
		}else{
      fmt.Print("\nNenhuma informação para planta cadastrada.")
    }
	}
}

func DeterminingComissionByPlantName(searchedName string) {

	for _, p := range plants {
		if p.Name == searchedName {
      fmt.Println("Comissão do vendedor: ", p.ComissionSeller())
		}else {
      fmt.Println("Comissão não definida.")
    }
	}
}
