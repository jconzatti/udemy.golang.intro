package main

import "fmt"

func multiblicacao(a, b int) int {
	return a * b
}

func soma(a, b int) int {
	return a + b
}

func subtracao(a, b int) int {
	return a - b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func main() {
	fmt.Printf("%d*%d=%d\n", 3, 5, exec(multiblicacao, 3, 5))
	fmt.Printf("%d+%d=%d\n", 3, 5, exec(soma, 3, 5))
	fmt.Printf("%d-%d=%d\n", 3, 5, exec(subtracao, 3, 5))
}
