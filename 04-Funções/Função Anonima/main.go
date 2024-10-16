package main

import "fmt"

func main() {
	func() {
		fmt.Println("Olá Mundo 1!")
	}()

	func(texto1 string) {
		fmt.Println(texto1)
	}("Olá Mundo 2!")

	retornoTexto := func(texto2 string) string {
		return fmt.Sprint("Olá ", texto2)
	}("Mundo 3!")

	fmt.Println(retornoTexto)
}
