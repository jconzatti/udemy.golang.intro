package main

import "fmt"

func main() {
	i := 1
	var p *int = nil //ponteiro
	p = &i           //obtendo endereço de i
	*p++             //obtém valor associado ao ponteiro
	i++
	fmt.Println(p, *p, &i, i)

	//não é permitido aritimética de ponteiros
	//p++ //isso não é permitido

}
