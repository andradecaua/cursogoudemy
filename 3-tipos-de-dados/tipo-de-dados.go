package main

import "fmt"

func main() {
	var numero int64 = -10000000000000
	fmt.Println(numero)

	// UINT aceita apenas números sem sinais
	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//alias = apelidos para alguns tipos
	//INT 32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	//BYTE = UINT8 == 8 bytes
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 132.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1328888.45
	fmt.Println(numeroReal2)

	numeroReal3 := 12333.67
	fmt.Println(numeroReal3)

	//STRINGS
	var str string = "asdadaudadha"
	fmt.Println(str)

	str2 := "texto2"
	fmt.Println(str2)

	char := '8'
	fmt.Println(char)
	// fim strings

	var texto int
	fmt.Println(texto)
}
