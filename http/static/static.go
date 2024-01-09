package main

import (
	"log"
	"net/http"
)

func main() {
	servidorDeArquivo := http.FileServer(http.Dir("public"))
	http.Handle("/", servidorDeArquivo)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
