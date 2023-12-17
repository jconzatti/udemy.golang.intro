package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//maps devem ser inicializados assim
	//var aprovados map[int]string = make(map[int]string)
	//ou assim
	aprovados := make(map[int]string)
	aprovados[2023120] = "Thaís"
	aprovados[2023121] = "Jhoni"
	aprovados[2023122] = "Shelby"
	fmt.Println(aprovados)

	for matricula, nome := range aprovados {
		fmt.Printf("%s (Matrícula: %d)\n", nome, matricula)
	}
	fmt.Println(aprovados[2023120])
	delete(aprovados, 2023121)
	fmt.Println(aprovados[2023121])
}
