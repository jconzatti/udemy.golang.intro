package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	fmt.Println("Teste " + string(97)) //converte para o valor mapeano na tabela unicode

	//IntToStr
	fmt.Println("Teste " + strconv.Itoa(97))

	//StrToInt
	num, _ := strconv.Atoi("97")
	fmt.Println(num - 96)

	//StrToBool
	b, _ := strconv.ParseBool("true") //"1" tbm é true
	if b {
		fmt.Println("Verdadeiro")
	}
}
