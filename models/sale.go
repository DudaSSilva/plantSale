package models

type Sale struct {
	Plant  Plant
	Seller Seller
	Client Client
	plants []Plant
}
