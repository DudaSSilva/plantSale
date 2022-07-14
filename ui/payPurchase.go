package ui

import (
	"errors"
	"fmt"
	"main/data"
)

func FinishPayment() float32 {
	var plantCode, quantPlants, i int
	var total float32 = 0

	fmt.Print("Quantas plantas estão sendo vendidas? ")
	fmt.Scan(&quantPlants)

	for i = 0; i < quantPlants; i++ {
		fmt.Print("Por favor, informe novamente o código da planta: ")
		fmt.Scan(&plantCode)

		plant, errorP := data.SearchPlantByName(plantCode)

		if errorP == nil {
			total += plant.Price
		} else {
			fmt.Print("\nErro! ", errorP)
		}

	}

	return total

}

func PayPurchase(total float32) (float32, error) {
	var paymentAmount, value float32

	fmt.Print("\nValor pago pelo cliente: R$ ")
	fmt.Scan(&value)

	if value >= total {
		paymentAmount = value - total
		return paymentAmount, nil
		fmt.Printf("\n\n %f", paymentAmount)
	}

	return value, errors.New("Saldo insuficiente.")

}
