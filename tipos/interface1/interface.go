package main

import "fmt"

type imprimivel interface {
	toString() string
}

type pessoa struct {
	nome      string
	sobrenome string
}

type produto struct {
	descricao string
	preco     float64
}

func (p pessoa) toString() string {
	return fmt.Sprintf("%s %s", p.nome, p.sobrenome)
}

func (p produto) toString() string {
	return fmt.Sprintf("%s R$%.2f", p.descricao, p.preco)
}

func imprimir(imp imprimivel) {
	fmt.Println(imp.toString())
}

func main() {
	var coisa imprimivel = pessoa{nome: "Jhoni", sobrenome: "Conzatti"}
	imprimir(coisa)
	coisa = produto{descricao: "Café", preco: 17.89}
	imprimir(coisa)
	p1 := produto{"Calça Jeans", 179.90}
	imprimir(p1)
}
