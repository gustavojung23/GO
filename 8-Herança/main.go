package main

import "fmt"

type Animal struct {
	Nome string
}

type Cachorro struct {
	Animal
	Raça string
}

func main() {
	c := Cachorro{
		Animal: Animal{Nome: "Pingo"},
		Raça:   "Vira Lata",
	}

	c.Falar()
	c.Animal.Falar()
}

func (c Cachorro) Falar() {
	fmt.Println(c.Nome, "diz: Au au!")
}

func (a Animal) Falar() {
	fmt.Println(a.Nome, "emite um som.")
}
