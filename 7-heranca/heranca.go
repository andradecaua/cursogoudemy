package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float64
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Cauã", "Pedro", 19, 178}
	fmt.Println(p1)
	e1 := estudante{p1, "Ciencia da computaco", "unifacvest"}
	fmt.Println(e1)
}
