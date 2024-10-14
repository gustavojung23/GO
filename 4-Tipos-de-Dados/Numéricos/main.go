package main

import "fmt"

func main() {
	numero := 100000
	fmt.Println("Numero int: ", numero)

	var numeroInt8 int8 = 127
	fmt.Println("Numero int8: ", numeroInt8)

	var numeroInt16 int16 = 32767
	fmt.Println("Número int16: ", numeroInt16)

	var numeroInt32 int32 = 2147483647
	fmt.Println("Número int32: ", numeroInt32)

	var numeroInt64 int64 = 9223372036854775807
	fmt.Println("Número int64: ", numeroInt64)

	var numeroUint uint = 2147483647
	fmt.Println("Número uint: ", numeroUint)

	var numeroUint8 uint8 = 255
	fmt.Println("Número uint8: ", numeroUint8)

	var numeroUint16 uint16 = 65535
	fmt.Println("Número uint16: ", numeroUint16)

	var numeroUint32 uint32 = 4294967295
	fmt.Println("Número uint32: ", numeroUint32)

	var numeroUint64 uint64 = 18446744073709551615
	fmt.Println("Número uint64: ", numeroUint64)

	var numeroRune rune = 2147483647
	fmt.Println("Número rune: ", numeroRune)

	var numeroByte byte = 255
	fmt.Println("Número byte: ", numeroByte)

	var numeroFloat32 float32 = 1.12345678901234567890
	fmt.Println("Número float64: ", numeroFloat32)

	var numeroFloat64 float64 = 1.12345678901234567890
	fmt.Println("Número float64: ", numeroFloat64)
}
