package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(5, 8))
	var subtracao = func(a, b int) int {
		return a - b
	}
	fmt.Println(subtracao(5, 8))
}
