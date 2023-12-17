package main

import "fmt"

func main() {
	x := 1
	y := 2
	x++ //x += 1, x = x + 1
	y--
	fmt.Println(x, y)

	//isso não pode em go: ainda bem!!!
	//fmt.Println(x == y--)
}
