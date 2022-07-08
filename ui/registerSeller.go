package ui

import (
	"main/data"
	"main/models"
	"fmt"
)

func RegisterSeller() {
	var name string
	var sellerCode int

	fmt.Print("\nVendedor responsável: ")
	fmt.Scan(&name)
	fmt.Print("Código do vendedor: ")
	fmt.Scan(&sellerCode)

	seller := models.Seller{
		Name:        name,
		SellerCode: sellerCode,
	}
  
	data.SaveSeller(seller)
}