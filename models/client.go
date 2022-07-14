package models

import "fmt"

type Client struct {
	Name string
	Cpf  string
}

func NewClient(name string) *Client {
	return &Client{Name: name}
}

func (c Client) ViewClientInformation() {
	fmt.Println("\nNome do cliente: ", c.Name, "\nCpf do cliente: ", c.Cpf)
}
