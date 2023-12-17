package main

import (
	"fmt"
	"reflect"
)

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 float64) {
	fmt.Printf("F2: %s %.2f\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, float64) {
	return "jhoni.conzatti@gmail.com", 8500
}

func main() {
	f1()
	f2("Parâmetro 1", 1)
	r3, r4 := f3(), f4("Parâmetro 1", "Parâmetro 2")
	fmt.Println(r3)
	fmt.Println(r4)
	r51, r52 := f5()
	fmt.Println("F5:", r51, r52, reflect.TypeOf(r52))
}
