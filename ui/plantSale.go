package ui

import (
	"fmt"
	"main/data"
)

func PlantSale() {
	var quantPlants, i int
	var answer string

	fmt.Print("\nInforme a quantidade de plantas a ser vendida: ")
	fmt.Scan(&quantPlants)

	for i = 0; i < quantPlants; i++ {
		RegisterPlant()
		RegisterSeller()
		RegisterClient()
	}

	fmt.Println("\nDeseja exibir informações e finalizar venda? ")
	fmt.Scan(&answer)

	if answer == "sim" {
		SellPlant()
	}
}

func SellPlant() {

	var plantCode int
	var cpf string
	var sellerCode int
	var finish string

	fmt.Print("\nCódigo da planta: ")
	fmt.Scan(&plantCode)
	fmt.Print("Código do vendedor: ")
	fmt.Scan(&sellerCode)
	fmt.Print("Cpf do cliente: ")
	fmt.Scan(&cpf)

	plant, errorP := data.SearchPlantByName(plantCode)
	seller, errorS := data.SearchSellerByCode(sellerCode)
	client, errorC := data.SearchClientByCpf(cpf)

	if errorP == nil && errorS == nil && errorC == nil {
		plant.ViewPlantInformation()
		seller.ViewSellerInformation()
		saleComission, _ := ComissionSeller(*plant)
		fmt.Println("Comissão do vendedor: R$ ", saleComission)
		client.ViewClientInformation()
	} else {
		fmt.Print("\nErro! ", errorP)
		fmt.Print("\nErro! ", errorS)
		fmt.Print("\nErro! ", errorC)
	}

	fmt.Print("\nDigite sim para finlizar a compra: ")
	fmt.Scan(&finish)

	if finish == "sim" {
		transshipment, _ := PayPurchase(FinishPayment())
		fmt.Printf("Valor do troco: R$ %.2f", transshipment)
	}

	fmt.Println("\n\nVENDA FINALIZADA!")
}
