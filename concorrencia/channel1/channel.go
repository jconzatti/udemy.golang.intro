package main

import "fmt"

func main() {
	canal := make(chan int, 1)
	canal <- 1 //escrevendo o valor em um canal
	<-canal    //lendo o valor de um canal

	canal <- 2
	fmt.Println(<-canal)
}
