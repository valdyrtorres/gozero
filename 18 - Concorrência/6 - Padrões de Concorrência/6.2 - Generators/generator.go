package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Mundo!")
	// Note que esse canal somente recebe
	// Ao utilizar o padrão, não foi necessário chamar como go routine aqui no 'main' ou controlar com o make
	// O padrão generator esconde a complexidade ao encapsular a chamadas na função escrever

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal

}
