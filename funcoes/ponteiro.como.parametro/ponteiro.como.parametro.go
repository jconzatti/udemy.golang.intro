package main

import "fmt"

func incrementarV1(n int) {
	n++
}

func incrementarV2(n *int) {
	*n++
}

func main() {
	n := 10
	incrementarV1(n)
	fmt.Println("n =", n)
	incrementarV2(&n)
	fmt.Println("n =", n)
}
