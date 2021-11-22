package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// waitgroup é um meio que eu tenho para sincronizar goroutines
	// Não é tão utilizado
	// Os canais são mais utilizados e é um jeito mais sofisticado
	var waitGroup sync.WaitGroup

	waitGroup.Add(4) // quantidade de goroutines que vão rodar ao mesmo tempo no grupo de espera

	// função anonima. O go na frente indica que a função anonima vai rodar de modo concorrente
	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // -1 retira um da fila
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("GoRoutine 3!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("GoRoutine 4!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
