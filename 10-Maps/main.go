package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Jhon",
		"sobrenome": "Doe",
	}
	fmt.Println(usuario)

	curso := map[string]map[string]string{
		"Faculdade": {
			"Campus": "UNI",
			"Curso":  "ADS",
		},
	}
	fmt.Println(curso)
}
