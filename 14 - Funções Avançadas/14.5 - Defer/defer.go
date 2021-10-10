package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Mádia calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		//fmt.Println("Mádia calculada. Resultado será retornado!")
		return true
	}

	//fmt.Println("Mádia calculada. Resultado será retornado!")
	return false
}

func main() {
	//defer funcao1() // DEFER = ADIAR
	//funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}
