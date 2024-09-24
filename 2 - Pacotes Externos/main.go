package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello World")

	erro := checkmail.ValidateFormat("teste@teste.com")
	fmt.Println(erro)

	erro2 := checkmail.ValidateFormat("123")
	fmt.Print(erro2)
}
