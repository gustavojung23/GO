package main

import "fmt"

func main() {
	soma, subtracao := calculos(10, 20)
	fmt.Println(soma, subtracao)
}

func calculos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}
