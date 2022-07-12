package ui

import (
	"fmt"
	"main/data"
	"main/models"
)

func PlantSale() {

	var name string
	var cpf string
	var sellerCode int

	fmt.Print("\nNome da planta: ")
	fmt.Scan(&name)
	fmt.Print("Código do vendedor: ")
	fmt.Scan(&sellerCode)
	fmt.Print("Cpf do cliente: ")
	fmt.Scan(&cpf)

	data.SearchPlantByName(name)
	data.ViewPlant(name)
	data.SearchSellerByCode(sellerCode)
	data.ViewSeller(sellerCode)
	data.SearchClientByCpf(cpf)
	data.ViewClient(cpf)

	sale := models.Sale{
		Name:       name,
		SellerCode: sellerCode,
		Cpf:        cpf,
	}

	plant, errorP := data.SearchPlantByName(name)
	//_, errorS := data.SearchSellerByCode(sellerCode)
	//_, errorC := data.SearchClientByCpf(cpf)
	//
	//if errorP != nil && errorS != nil && errorC != nil {
	//	data.ViewPlant(name)
	//	data.ViewSeller(sellerCode)
	//	data.ViewClient(cpf)
	//}

	if errorP != nil {
		sale.ViewComissionSeller(plant)
	}

}

func SellPlant() {
	var quantPlants, i int

	fmt.Print("\nInforme a quantidade de plantas a ser vendida: ")
	fmt.Scan(&quantPlants)

	for i = 0; i < quantPlants; i++ {
		RegisterPlant()
		RegisterSeller()
		RegisterClient()
		fmt.Println("\nVenda realizada!")
	}
}
