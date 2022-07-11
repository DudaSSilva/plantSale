package main

import (
	"fmt"
	"main/ui"
)

func main() {
	var answer int
	var turn = "sim"

	for turn != "nao" {

		fmt.Println("\n======== MENU ========")
		fmt.Println("\n1 - Vender planta\n2 - Exibir informações de venda")
		fmt.Print("\nSelecione uma das opções acima: ")
		fmt.Scan(&answer)

		if answer == 1 {
			ui.SellPlant()
		} else if answer == 2 {
			ui.PlantSale()
		} else {
			fmt.Println("\nOpção inválida")
		}

		fmt.Println("\nVoltar para o menu?")
		fmt.Scan(&turn)
	}

}
