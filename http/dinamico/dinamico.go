package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func horaCerta(w http.ResponseWriter, r *http.Request) {
	horaFormatada := time.Now().Format("02/01/2006 15:04:05")
	fmt.Fprintf(w, "<h1>Hora certa: %s</h1>", horaFormatada)
}

func main() {
	http.HandleFunc("/horaCerta", horaCerta)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
