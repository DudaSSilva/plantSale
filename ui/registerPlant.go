package ui

import (
	"fmt"
	"main/data"
	"main/models"
)

func RegisterPlant() {
	var name, species, material, properties string
	var price float32
	var inventory int
	var plantCode int

	fmt.Print("\nNome da planta: ")
	fmt.Scan(&name)
	fmt.Print("Código da planta: ")
	fmt.Scan(&plantCode)
	fmt.Print("Espécie: ")
	fmt.Scan(&species)
	fmt.Print("Material: ")
	fmt.Scan(&material)
	fmt.Print("Planta(s) em estoque: ")
	fmt.Scan(&inventory)
	fmt.Print("Valor: R$ ")
	fmt.Scan(&price)
	fmt.Print("Propriedades da planta: ")
	fmt.Scan(&properties)

	plant := models.Plant{
		Name:       name,
		PlantCode:  plantCode,
		Species:    species,
		Material:   material,
		Inventory:  inventory,
		Price:      price,
		Properties: properties,
	}

	data.SavePlant(plant)
}
