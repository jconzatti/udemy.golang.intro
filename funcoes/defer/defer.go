package main

import "fmt"

func obterValorAprovado(numero int) int {
	defer fmt.Println("Imprime logo antes de cada return 1")
	defer fmt.Println("Imprime logo antes de cada return 2")
	if numero >= 5000 {
		defer fmt.Printf("Valor %d muito alto!\n", numero)
		return 5000
	}
	defer fmt.Printf("Valor %d aceito!\n", numero)
	return numero
}

func main() {
	fmt.Println("Resultado 1:", obterValorAprovado(6000))
	fmt.Println("Resultado 2:", obterValorAprovado(456))
}
