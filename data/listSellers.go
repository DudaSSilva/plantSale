package data

import (
	"errors"
	"main/models"
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
