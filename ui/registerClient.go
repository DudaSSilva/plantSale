package ui

import (
	"fmt"
	"main/data"
	"main/models"
)

func RegisterClient() {
	var name, cpf string

	fmt.Print("\nNome do cliente: ")
	fmt.Scan(&name)
	fmt.Print("Cpf do cliente: ")
	fmt.Scan(&cpf)

	client := models.Client{
		Name:     name,
		Cpf: cpf,
	}

	data.SaveClient(client)
}