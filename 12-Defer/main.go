package main

import "fmt"

func main() {
	fmt.Println(alunoAprovado(7, 6))
}

func alunoAprovado(n1, n2 float32) string {
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Calculando Média.")

	media := (n1 + n2) / 2

	if media >= 6 {
		return fmt.Sprint("Aluno aprovado!")
	}
	return fmt.Sprintf("Aluno reprovado!")
}
