package main

import "fmt"

func main() {
	retornoAnonimo := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Parametro")

	fmt.Println(retornoAnonimo)
}
