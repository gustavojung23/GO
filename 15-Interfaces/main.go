package main

import (
	"fmt"
	"math"
)

type Area interface {
	area() float64
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

type Circulo struct {
	Raio float64
}

func escreverArea(a Area) {
	fmt.Printf("A forma da área é: %0.2f\n", a.area())
}

func (r Retangulo) area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}

func main() {
	r := Retangulo{5, 20}
	escreverArea(r)

	c := Circulo{5}
	escreverArea(c)
}
