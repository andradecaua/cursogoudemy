package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	/*
		for i < 10 {
			i++
			fmt.Println("Incrementando i")
			time.Sleep(time.Second)
		} */

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Printf("Incrementando j = %d\n", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Cauã", "luiza", "helena"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Souza",
	}

	fmt.Println(usuario)

	for key, valor := range usuario {
		fmt.Println(key, valor)
	}

}
