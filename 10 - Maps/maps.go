package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[int]string{
		1: "Pedro",
		2: "Silva",
	}

	fmt.Println(usuario2)
	fmt.Println(usuario2[1])

	// aninhado
	usuario3 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},
		"curso": {
			"nome":   "Engenharia",
			"ultimo": "Campus 1",
		},
	}

	fmt.Println(usuario3)

	// deletando uma chave
	delete(usuario3, "nome")
	fmt.Println(usuario3)

	usuario3["signo"] = map[string]string{
		"nome": "Gêmeos",
	}

	fmt.Println(usuario3)
}
