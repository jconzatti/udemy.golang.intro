package main

import (
	"fmt"
	"time"
)

func rotina(canal chan int) {
	canal <- 1
	canal <- 2
	canal <- 3
	canal <- 4
	fmt.Println("Executou!")
	canal <- 5
	canal <- 6
}

func main() {
	canal := make(chan int, 3)
	go rotina(canal)

	time.Sleep(time.Second)
	fmt.Println(<-canal)
}
