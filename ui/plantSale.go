package ui

import (
	"fmt"
	"main/data"
)

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
  
  data.SearchPlantByName(name)
  data.ViewPlant(name)
  data.SearchSellerByCode(sellerCode)
  data.ViewSeller(sellerCode)
  data.DeterminingComissionByPlantName(name)
  data.SearchClientByCpf(cpf)
  data.ViewClient(cpf)
}

func SellPlant(){
  var quantPlants, i int
  
  fmt.Print("\nInforme a quantidade de plantas a ser vendida: ")
  fmt.Scan(&quantPlants)
  for i = 0; i < quantPlants; i++ {
    RegisterPlant()
    RegisterSeller()
    RegisterClient()
    fmt.Println("\nVenda realizada!")
  }
}
