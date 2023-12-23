package main

import "fmt"

type item struct {
	codigoDoProduto int
	quantidade      float64
	precoUnitario   float64
}

type pedido struct {
	codigoDoUsuario int
	itens           []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.quantidade * item.precoUnitario
	}
	return total
}

func main() {
	pedido1 := pedido{
		codigoDoUsuario: 1,
		itens: []item{
			{1, 2, 3},
			{2, 3, 4},
			{3, 4, 5},
			{4, 1, 2},
		},
	}
	fmt.Printf("O valor total do pedido é R$ %.2f\n", pedido1.valorTotal())
}
