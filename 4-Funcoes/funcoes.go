package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	var n1 int8 = 10
	var n2 int8 = 10

	fmt.Printf("A soma de %d e %d = %d\n", n1, n2, somar(n1, n2))

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	resultadoSoma, _ := calculosMatematicos(n1, n2)

	fmt.Println(resultadoSoma)
}
