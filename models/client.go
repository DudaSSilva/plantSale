package models

import "fmt"

type Client struct {
	Name string
	Cpf  string
}

func (c Client) ViewClientInformation() {
	fmt.Println("\nNome do cliente: ", c.Name, "\nCpf do cliente: ", c.Cpf)
}
