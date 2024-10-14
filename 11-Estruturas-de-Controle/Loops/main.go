package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println("Laço I: ", i)
	}

	cidades := [3]string{"Nova York", "Brasília", "Paris"}
	for i, cidade := range cidades {
		fmt.Println(i, cidade)
	}
}
