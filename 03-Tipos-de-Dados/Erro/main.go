package main

import (
	"errors"
	"fmt"
)

func main() {
	var erro error
	fmt.Println(erro)

	var newErro error = errors.New("Erro interno")
	fmt.Println(newErro)
}
