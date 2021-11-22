package main

import (
	"fmt"
	"time"
)

// o canal pode tanto enviar como receber dados

func main() {
	canal := make(chan string) //indica que o canal só pode enviar e receber strings

	go escrever("Olá Mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem, aberto := <-canal // O canal espera as entradas para enviar e a variável mensagem armazena esse valor
		// a variavel aberto server para saber se o canal está aberto ou fechado
		if !aberto {
			break
		}

		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // Aqui envio o dado ao canal
		time.Sleep(time.Second)
	}

	close(canal)
}
