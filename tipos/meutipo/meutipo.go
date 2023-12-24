package main

import "fmt"

type nota float64

func (n nota) entre(notaInicial, notaFinal nota) bool {
	return n >= notaInicial && n <= notaFinal
}

func (n nota) conceito() string {
	switch {
	case n.entre(9, 10):
		return "A"
	case n.entre(8, 9):
		return "B"
	case n.entre(7, 8):
		return "C"
	case n.entre(5, 7):
		return "D"
	case n.entre(3, 5):
		return "E"
	default:
		return "F"
	}
}

func (n nota) imprimirConceito() {
	fmt.Printf("Conceito nota %.2f: %s\n", n, n.conceito())
}

func main() {
	nota(9.8).imprimirConceito()
	nota(6.9).imprimirConceito()
	nota(3.2).imprimirConceito()
}
