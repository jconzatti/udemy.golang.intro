package main

import "fmt"

func trocar(p1, p2 int) (segundo int, primeiro int) {
	primeiro = p1
	segundo = p2
	return //retorno limpo: as variáveis de retorno já foram atribuidas acima
}

func main() {
	x, y := trocar(2, 3)
	fmt.Println(x, y)
}
