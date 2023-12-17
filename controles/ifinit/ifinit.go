package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { //tbm permite no case
		fmt.Println(i, "Ganhou :D")
	} else {
		fmt.Println(i, "Perdeu :(")
	}
}
