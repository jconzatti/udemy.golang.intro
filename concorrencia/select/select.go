package main

import (
	"curso/concorrencia/html"
	"fmt"
	"time"
)

func oMaisRapido(url1, url2, url3 string) string {
	canal1 := html.Titulos(url1)
	canal2 := html.Titulos(url2)
	canal3 := html.Titulos(url3)

	select {
	case titulo1 := <-canal1:
		return titulo1
	case titulo2 := <-canal2:
		return titulo2
	case titulo3 := <-canal3:
		return titulo3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		//default:
		//	return "Sem resposta ainda!"
	}
}

func main() {
	campeao := oMaisRapido("https://www.chess.com", "https://www.alura.com.br", "https://www.google.com")
	fmt.Println(campeao)
}
