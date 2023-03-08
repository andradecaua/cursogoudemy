package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradadouro string
	numero       uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Cauã"
	u.idade = 19

	enderecoExemplo := endereco{"Rua Caiana", 21}
	u2 := usuario{"Caua", 21, enderecoExemplo}

	u3 := usuario{nome: "Luiza"}

	fmt.Println(u)
	fmt.Println(u2)
	fmt.Println(u3)
}
