package main

import (
	"curso/concorrencia/html"
	"fmt"
)

func encaminhar(canalDeOrigem <-chan string, canalDeDestino chan string) {
	for {
		canalDeDestino <- <-canalDeOrigem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	canal := make(chan string)
	go encaminhar(entrada1, canal)
	go encaminhar(entrada2, canal)
	return canal
}

func main() {
	canal := juntar(
		html.Titulos("https://www.chess.com", "https://www.alura.com.br"),
		html.Titulos("https://www.google.com", "https://www.youtube.com"),
	)
	fmt.Println(<-canal, "|", <-canal)
	fmt.Println(<-canal, "|", <-canal)
}
