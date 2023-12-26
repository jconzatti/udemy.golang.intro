package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	realizarBalisa()
}

type esportivoLuxuoso interface {
	esportivo
	luxuoso
}

type bmw7 struct{}

func (b bmw7) ligarTurbo() {
	fmt.Println("Turbo ligado!")
}

func (b bmw7) realizarBalisa() {
	fmt.Println("Realizando balisa!")
}

func main() {
	var b esportivoLuxuoso = bmw7{}
	b.ligarTurbo()
	b.realizarBalisa()
}
