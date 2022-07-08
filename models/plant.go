package models

import "fmt"

type Plant struct {
	Name string
  PlantCode int
	Species string
	Material     string
  Inventory int
	Price        float64
	Properties string
}

func (plant Plant) ComissionSeller() float64 {
	return (5 / 100) * plant.Price
}


func (p Plant) ViewPlantInformation() {
	fmt.Println("\nNome da planta: ", p.Name, "\nCódigo da planta: ", p.PlantCode, "\nEspécie: ", p.Species, "\nMaterial: ", p.Material,
              "\nEstoque: ", p.Inventory, "\nPreço: ", p.Price, "\nPropriedades: ", p.Properties)
}