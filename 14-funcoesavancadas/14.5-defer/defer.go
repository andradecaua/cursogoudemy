package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média Calculada. Resultado sera retornado")
	fmt.Println("Entrando na função para verificar se o aluno foi aprovado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}

func main() {
	resultadoAlunoAprovado := alunoEstaAprovado(7.2, 5.7)
	fmt.Println(resultadoAlunoAprovado)
}
