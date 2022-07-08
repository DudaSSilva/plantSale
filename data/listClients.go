package data

import (
  "fmt"
	"main/models"
	"errors"
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

func ViewClient(clientCpf string){

	for _, c := range clients {
		if c.Cpf == clientCpf {
			c.ViewClientInformation()
		}else{
      fmt.Print("\nNenhuma informação para cliente cadastrada.")
    }
	}

}