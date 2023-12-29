package html

import (
	"io"
	"net/http"
	"regexp"
)

// Titulos obtém os títulos das páginas de um conjunto de URLs
func Titulos(urls ...string) <-chan string {
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
