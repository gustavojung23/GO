package main

import "fmt"

func main() {
	numero1 := 10

	if numero1 >= 15 {
		fmt.Println("Maior")
	} else {
		fmt.Println("Menor")
	}

	if numero2 := numero1; numero2 > 9 {
		fmt.Println("Maior")
	} else {
		fmt.Println("Menor")
	}
}
