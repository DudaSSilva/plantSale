package models

type Sale struct {
	Name       Plant
	SellerCode Seller
	Cpf        Client
}

func (plant Plant) ComissionSeller() float64 {
	return (5 / 100) * plant.Price
}

//func (plant Plant) ViewComissionSeller() {
//	fmt.Println("\nComissão do vendedor: ", plant.ComissionSeller())
//}
