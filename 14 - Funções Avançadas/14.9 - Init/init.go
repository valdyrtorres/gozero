package main

import "fmt"

var n int

// você pode ter uma função init por arquivo
func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Função main sendo executada")
	fmt.Println(n)
}
