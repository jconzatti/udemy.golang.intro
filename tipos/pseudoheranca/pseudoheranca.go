package main

import "fmt"

type carro struct {
	nome            string
	velocidadeAtual int
}

type ferrari struct {
	carro
	turboLigado bool
}

func main() {
	f := ferrari{}
	f.nome = "Ferrari F40"
	f.velocidadeAtual = 100
	f.turboLigado = true
	fmt.Printf("O carro %s -> velocidade: %d, turbo: %t\n", f.nome, f.velocidadeAtual, f.turboLigado)
	fmt.Println(f)
}
