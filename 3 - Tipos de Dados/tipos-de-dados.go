package main

import (
	"errors"
	"fmt"
)

func main() {
	var nunero int64 = 100000000000000
	fmt.Println(nunero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	// eh um alias
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE
	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1230000000.45
	fmt.Println(numeroReal3)

	// FIM NÚMEROS REAIS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	// FIM STRINGS

	texto := 5
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
