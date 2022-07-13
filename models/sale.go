package models

type Sale struct {
	Name       string
	SellerCode int
	Cpf        string
	//Price      float32
}

//func (s Sale) ComissionSeller(price float32) float32 {
//	return ((5 / 100) * price)
//}

//func (s Sale) ViewComissionSeller(price float32) {
//	fmt.Println("\nComissão do vendedor: ")
//	fmt.Printf("%f", s.ComissionSeller(price))
//	fmt.Print("\n\n\n")
//	fmt.Print(price)
//}
