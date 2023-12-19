package main

import "fmt"

func imprimirCandidatos(candidatos ...string) {
	fmt.Println("Lista de candidatos")
	for i, candidato := range candidatos {
		fmt.Printf("%d) %s\n", i+1, candidato)
	}
}

func main() {
	candidatos := []string{"Maria", "José", "Jhoni", "Thaís", "Tayná", "João"}
	imprimirCandidatos(candidatos...)
}
