package main

import "fmt"

func recuperarExecuca() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(medias ...float64) bool {
	defer recuperarExecuca()
	var total float64 = 0.0
	for _, media := range medias {
		total += media
	}

	total = total / float64(len(medias))

	if total > 6 {
		return true
	} else if total < 6 {
		return false
	}

	panic("A MEDIA É EXATAMENTE 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pos execução!")
}
