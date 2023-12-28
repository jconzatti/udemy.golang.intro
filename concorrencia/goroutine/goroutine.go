package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, quantidade int) {
	for i := 0; i < quantidade; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	//fale("Maria", "Porque você não falo comigo?", 3)
	//fale("Jõao", "Posso falar somente depois de você!", 1)

	//go fale("Maria", "Ei!!!", 500)
	//go fale("Jõao", "Opa!", 500)
	//time.Sleep(5 * time.Second)
	//fmt.Println("Fim!")

	go fale("Maria", "Entendi!", 50)
	fale("Jõao", "Parabéns!", 5)
}
