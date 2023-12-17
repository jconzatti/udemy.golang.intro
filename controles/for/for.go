package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i, "Jhoni Conzatti")
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d", i)
	}

	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Printf("%d é par", i)
		} else {
			fmt.Printf("%d é ímpar", i)
		}
	}

	for {
		fmt.Println("Para sempre...")
		time.Sleep(time.Second)
	}
}
