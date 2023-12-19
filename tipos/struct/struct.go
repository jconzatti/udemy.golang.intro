package main

import "fmt"

type produto struct {
	codigo               int
	descricao            string
	preco                float64
	percentualDeDesconto float64
}

func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.percentualDeDesconto)
}

func main() {
	var produto1 produto = produto{
		codigo:               1,
		descricao:            "Café",
		preco:                15.85,
		percentualDeDesconto: 0.1,
	}
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{2, "Arroz", 19.50, 0.05}
	produto2.percentualDeDesconto = 0.08
	fmt.Println(produto2, produto2.precoComDesconto())
}
