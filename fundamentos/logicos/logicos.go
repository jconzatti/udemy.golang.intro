package main

import "fmt"

func comprar(trab1, trab2 bool) (bool, bool, bool) {
	comprarTV50 := trab1 && trab2
	comprarTV32 := trab1 != trab2
	comprarSorvete := trab1 || trab2
	return comprarTV50, comprarTV32, comprarSorvete
}

func main() {
	tv50, tv32, sorvete := comprar(true, true)
	fmt.Printf("TV 50': %t, TV 32': %t, Sorvete: %t, Saudável: %t", tv50, tv32, sorvete, !sorvete)
}
