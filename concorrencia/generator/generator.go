package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

// <-chan: canal somente-leitura
func titulos(urls ...string) <-chan string {
	canalDeTitulos := make(chan string)
	for _, url := range urls {
		go func(url string) {
			respostaHTTP, _ := http.Get(url)
			html, _ := io.ReadAll(respostaHTTP.Body)
			r, _ := regexp.Compile("<title>(.*?)</title>")
			canalDeTitulos <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return canalDeTitulos
}

func main() {
	titulos1 := titulos("https://www.chess.com", "https://www.alura.com.br")
	titulos2 := titulos("https://www.google.com", "https://www.youtube.com")
	fmt.Println("Primeiros:", <-titulos1, "|", <-titulos2)
	fmt.Println("Segundos:", <-titulos1, "|", <-titulos2)
}
