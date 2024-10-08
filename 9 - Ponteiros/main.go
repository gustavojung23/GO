package main

import "fmt"

func main() {
	var number1 int
	var numberPointer *int

	number1 = 100
	numberPointer = &number1

	fmt.Println(number1, numberPointer)
}
