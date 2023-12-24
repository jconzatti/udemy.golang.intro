package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo     string
	turbo      bool
	velocidade float32
}

func (f *ferrari) ligarTurbo() {
	f.turbo = true
}

func main() {
	carro1 := ferrari{"F50", false, 0}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
