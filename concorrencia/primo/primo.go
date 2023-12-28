package main

import (
	"fmt"
	"time"
)

func numeroPrimo(numero int) bool {
	for i := 2; i < numero; i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true
}

func primos(quantidade int, canalDePrimos chan int) {
	inicio := 2
	for i := 0; i < quantidade; i++ {
		for p := inicio; ; p++ {
			if numeroPrimo(p) {
				canalDePrimos <- p
				inicio = p + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(canalDePrimos)
}

func main() {
	canalDePrimos := make(chan int, 30)
	go primos(100, canalDePrimos)
	for primo := range canalDePrimos {
		fmt.Printf("%d ", primo)
	}
	fmt.Print("Fim!")
}
