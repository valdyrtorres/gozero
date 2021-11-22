package main

import (
	"fmt"
	"time"
)

func main() {
	// CONCORRÊNCIA != PARALELISMO
	go escrever("Olá Mundo!") // go routine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
