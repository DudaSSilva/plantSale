package ui

import (
	"errors"
	"fmt"
	"main/models"
)

//type Payment struct{
//	PaymentAmount float64
//	Transshipment float64
//}

func PayPurchase(plant models.Plant) (float32, error) {
	var paymentAmount, value float32

	fmt.Print("\nValor pago pelo cliente: ")
	fmt.Scan(&value)

	if value >= plant.Price {
		paymentAmount = value - plant.Price
		return paymentAmount, nil
		fmt.Printf("\n\n %f", paymentAmount)
	}

	return value, errors.New("Saldo insuficiente.")

}
