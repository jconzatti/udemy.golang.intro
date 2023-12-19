package main

import "fmt"

func fatorial(n uint) uint {
	if n > 0 {
		return n * fatorial(n-1)
	}
	return 1
}

func main() {
	resultado := fatorial(50)
	fmt.Println("50! =", resultado)
}
