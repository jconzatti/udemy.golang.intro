package main

import (
	"fmt"
	"strings"
)

type pessoa struct {
	nome      string
	sobrenome string
}

func (p pessoa) nomeCompleto() string {
	return p.nome + " " + p.sobrenome
}

func (p *pessoa) atribuirNomeCompleto(nomeESobrenome string) {
	nomes := strings.Split(nomeESobrenome, " ")
	p.nome = nomes[0]
	p.sobrenome = nomes[1]
}

func main() {
	pessoa1 := pessoa{"Jhoni", "Conzatti"}
	fmt.Println(pessoa1.nomeCompleto())
	pessoa1.atribuirNomeCompleto("Maria Pereira da Silva")
	fmt.Println(pessoa1.nomeCompleto())
}
