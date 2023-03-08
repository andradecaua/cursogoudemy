package main

import "fmt"

func invererSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1

}

func main() {
	numero := 20
	numeroInvertido := invererSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)

	fmt.Println("------------------------")
	fmt.Println(&numero)
	fmt.Println(&novoNumero)
}
