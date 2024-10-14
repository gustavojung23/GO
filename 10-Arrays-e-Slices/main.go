package main

import "fmt"

func main() {
	var array1 [5]int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)

	array2 := [2]string{"Posição 1", "Posição 2"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4}
	fmt.Println(array3)

	slice1 := []int{10, 11, 12, 13}
	fmt.Println(slice1)

	slice1 = append(slice1, 14)
	fmt.Println(slice1)

	slice2 := slice1[0:3]
	fmt.Println(slice2)

	slice3 := make([]int, 3)
	fmt.Println("Tamanho do Slice: ", len(slice3))
	fmt.Println("Capacidade do Slice: ", cap(slice3))
}
