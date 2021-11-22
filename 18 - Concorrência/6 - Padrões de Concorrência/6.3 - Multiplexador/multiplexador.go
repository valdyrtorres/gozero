package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Esse padrão centraliza os canais em um só
	canal := multiplexar(escrever("Olá Mundo!"), escrever("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

// A função abaixo recebe dois canais de entrada e retorna um canal de saída
func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			//time.Sleep(time.Millisecond * 500)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
