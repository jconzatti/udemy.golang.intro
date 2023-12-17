package main

import "fmt"

func main() {
	funcionarios := map[string]float64{
		"Jhoni Conzatti":      5200.00,
		"Thaís de Lima Silva": 2800.00,
		"Tayná de Lima Silva": 3900.00,
	}
	funcionarios["Jhoni Conzatti"] = 8500.00
	delete(funcionarios, "Inexistente")
	for nome, salario := range funcionarios {
		fmt.Printf("%s (Salário: R$%.2f)\n", nome, salario)
	}
}
