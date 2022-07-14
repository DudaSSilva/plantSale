package main

import (
	"fmt"
	"main/ui"
)

func main() {
	var turn = "sim"

	for turn != "nao" {

		fmt.Println("\n========== VENDA DE PLANTA ==========")
		ui.PlantSale()

		fmt.Println("\nVender outra planta?")
		fmt.Scan(&turn)
	}

}
