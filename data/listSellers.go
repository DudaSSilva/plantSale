package data

import (
  "fmt"
	"main/models"
	"errors"
)

var sellers []models.Seller

func SearchSellerByCode(sellerCode int) (*models.Seller, error) {

	for _, v := range sellers {
		if v.SellerCode == sellerCode {
			return &v, nil
		}
	}

	return nil, errors.New("Vendedor não encontrado.")
}

func SaveSeller(seller models.Seller) []models.Seller {
	  sellers = append(sellers, seller)
	return sellers
}

func ViewSeller(sellerCode int){

	for _, v := range sellers {
		if v.SellerCode == sellerCode {
			v.ViewSellerInformation()
		}else{
      fmt.Print("\nNenhuma informação para vendedor cadastrada.")
    }
	}

}
