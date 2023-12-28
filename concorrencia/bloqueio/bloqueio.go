package main

import (
	"fmt"
	"time"
)

func rotina(canal chan int) {
	time.Sleep(time.Second)
	canal <- 1
	fmt.Println("Só depois que o canal foi lido")
}

func main() {
	canal := make(chan int)
	go rotina(canal)
	fmt.Println(<-canal)
	fmt.Println("Foi lido")
	//fmt.Println(<-canal)
	fmt.Println("Fim")
}
