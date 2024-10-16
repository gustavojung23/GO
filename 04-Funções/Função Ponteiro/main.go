package main

import "fmt"

func main() {
	number := 20
	inverterNumber(&number)
	fmt.Println(number)
}

func inverterNumber(number *int) {
	*number = *number * -1
}
