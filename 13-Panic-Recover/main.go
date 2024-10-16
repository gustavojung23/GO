package main

import "fmt"

func main() {
	fmt.Println(alunoAprovado(6, 6))
	fmt.Println("Outra função.")
}

func alunoAprovado(n1, n2 float32) string {
	defer recoveryFunc()
	fmt.Println("Calculando Média.")

	media := (n1 + n2) / 2

	if media > 6 {
		return fmt.Sprint("Aluno aprovado!")
	} else if media < 6 {
		return fmt.Sprintf("Aluno reprovado!")
	}

	panic("A média é 6!")
}

func recoveryFunc() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada!")
	}
}
