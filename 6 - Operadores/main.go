package main

import "fmt"

func main() {
	//operadores aritmeticos
	soma := 2 + 2
	subtracao := 2 - 1
	multiplicacao := 2 * 5
	divisao := 10 * 5
	restoDaDivisao := 10 % 2
	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDaDivisao)

	//operadores atribuição
	var variavel1 string = "String"
	variavel2 := "String 2"
	fmt.Println(variavel1, variavel2)

	//operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	//operadores lógicos
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//operadores unários
	numero1 := 10
	numero1++
	numero1 += 15
	fmt.Println(numero1)

	numero2 := 10
	numero2--
	numero2 -= 5
	fmt.Println(numero2)

	numero3 := 10
	numero3 *= 2
	fmt.Println(numero3)

	numero4 := 14
	numero4 /= 2
	fmt.Println(numero4)

	numero5 := 10
	numero5 %= 2
	fmt.Println(numero5)
}
