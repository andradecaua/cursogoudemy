package main

import "fmt"

func main() {
	//ARITIMETICOS

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restodaDivisao := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, restodaDivisao)

	var numer1 int16 = 10
	var numero2 int16 = 25
	soma2 := numer1 + numero2
	fmt.Println(soma2)

	// FIM DOS ARITIMETIOCS

	// ATRIBUICAO
	var variavel1 string = "string"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// FIM ATRIBUICAO

	//Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// Fim dos relacionais

	// Operadores Lógicos
	fmt.Println("--------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!falso)
	// Fim dos operadores logicos

	// Operadores unarios
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)
	numero--
	numero -= 10

	numero *= 2
	numero /= 2

	fmt.Println(numero)

}
