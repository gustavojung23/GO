package main

import "fmt"

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	//retornando mais de um valor na função
	resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao := calculos(4, 2)
	fmt.Println(resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao)

	//declarando uma variável como função
	var funcao = func(texto string) string {
		return texto
	}
	resultadoFuncao := funcao("Texto da Função")
	fmt.Println(resultadoFuncao)
}

func somar(numero1 int8, numero2 int8) int8 {
	return numero1 + numero2
}

// função pode retornar mais de um valor
func calculos(numero1, numero2 int8) (int8, int8, int8, int8) {
	soma := numero1 + numero2
	subtracao := numero1 - numero2
	multiplicacao := numero1 * numero2
	divisao := numero1 / numero2

	return soma, subtracao, multiplicacao, divisao
}
