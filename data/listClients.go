package data

import (
	"errors"
	"main/models"
)

var clients []models.Client

func SearchClientByCpf(clientCpf string) (*models.Client, error) {

	for _, c := range clients {
		if c.Cpf == clientCpf {
			return &c, nil
		}
	}

	return nil, errors.New("Cliente não encontrado.")
}

func SaveClient(client models.Client) []models.Client {
	clients = append(clients, client)
	return clients
}
