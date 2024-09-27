package main

import "fmt"

type Usuario struct {
	Nome     string
	Idade    uint8
	Endereco Endereco
}

type Endereco struct {
	Rua    string
	Numero uint8
}

func main() {
	var u Usuario
	u.Nome = "John"
	u.Idade = 18
	u.Endereco = Endereco{
		Rua:    "Rua Flor",
		Numero: 12,
	}

	fmt.Println(u)
}
