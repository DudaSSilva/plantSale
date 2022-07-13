package ui

import (
	"fmt"
	"main/data"
)

func SellPlant() {
	var quantPlants, i int

	fmt.Print("\nInforme a quantidade de plantas a ser vendida: ")
	fmt.Scan(&quantPlants)

	for i = 0; i < quantPlants; i++ {
		RegisterPlant()
		RegisterSeller()
		RegisterClient()
	}
}

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

	//sale := models.Sale{
	//	Name:       name,
	//	SellerCode: sellerCode,
	//	Cpf:        cpf,
	//}

	plant, errorP := data.SearchPlantByName(name)
	seller, errorS := data.SearchSellerByCode(sellerCode)
	client, errorC := data.SearchClientByCpf(cpf)

	if plant.Name == name && seller.SellerCode == sellerCode && client.Cpf == cpf {
		plant.ViewPlantInformation()
		seller.ViewSellerInformation()
		client.ViewClientInformation()
		saleComission, _ := ComissionSeller(*plant)
		fmt.Println("\nComissão do vendedor: ", saleComission)
		transshipment, _ := PayPurchase(*plant)
		fmt.Printf("\nValor do troco: %.2f", transshipment)
		fmt.Println("\nVenda finalizada!")

	} else {
		fmt.Print(errorP)
		fmt.Print(errorS)
		fmt.Print(errorC)
	}

}
