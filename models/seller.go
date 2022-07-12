package models

import "fmt"

type Seller struct {
	Name       string
	SellerCode int
	Comissao   float64
}

func (s Seller) ViewSellerInformation() {
	fmt.Println("\nNome do vendedor: ", s.Name, "\nCódigo do Vendedor: ", s.SellerCode)
}
