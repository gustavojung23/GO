package main

import "fmt"

func main() {
	dia := diaDaSemana(1)
	fmt.Println(dia)
}

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Segunda-Feira"
	case 4:
		return "Segunda-Feira"
	case 5:
		return "Segunda-Feira"
	case 6:
		return "Segunda-Feira"
	case 7:
		return "Segunda-Feira"
	default:
		return "Número Inválido"
	}
}
