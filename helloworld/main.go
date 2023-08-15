package main

import "fmt"

func main() {
	var v1 float64 = 0.512
	// Ponteiro para v1
	var v2 = &v1

	fmt.Println("Teste: ", v1, v2)

	// Alterando valor do conteúdo apontado pelo ponteiro
	*v2 = 1.64

	fmt.Println("Teste: ", v1, v2)
}
