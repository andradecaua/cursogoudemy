package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados", u.nome)
}

func main() {
	usuario1 := usuario{"Cauã", 20}
	fmt.Println(usuario1)
	usuario1.salvar()
}
