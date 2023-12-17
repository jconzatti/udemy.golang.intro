package main

import "fmt"

func notaParaConceito(nota float64) string {
	var parteInteiraDaNota = int(nota)
	switch parteInteiraDaNota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(1.2))
	fmt.Println(notaParaConceito(11.5))
}
