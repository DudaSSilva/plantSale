package models

import "fmt"

type Plant struct {
	Name       string
	PlantCode  int
	Species    string
	Material   string
	Inventory  int
	Price      float32
	Properties string
}

func (plant Plant) ViewPlantInformation() {
	fmt.Println("\nNome da planta: ", plant.Name, "\nCódigo da planta: ", plant.PlantCode, "\nEspécie: ", plant.Species, "\nMaterial: ", plant.Material,
		"\nEstoque: ", plant.Inventory, "\nPreço: ", plant.Price, "\nPropriedades: ", plant.Properties)
}
