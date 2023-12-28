package main

import (
	"fmt"
	"time"
)

//Channel (canal) - tipo chan - é o ponto de sincronismo das goroutines

func doisTresQuatroVezes(base int, canal chan int) {
	time.Sleep(time.Second)
	canal <- 2 * base

	time.Sleep(time.Second)
	canal <- 3 * base

	time.Sleep(time.Second)
	canal <- 4 * base
}

func main() {
	canal := make(chan int)
	go doisTresQuatroVezes(2, canal)
	fmt.Println("A")

	a, b := <-canal, <-canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println("C")
	fmt.Println(<-canal)
	//fmt.Println(<-canal) //deadlock
}
