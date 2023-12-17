package main

import "fmt"

func main() {
	funcionariosPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriel": 1200.00,
			"Guga":    1300.00,
		},
		"J": {
			"Jhoni": 8500,
		},
		"T": {
			"Thaís": 2800,
			"Tayná": 3900,
		},
		"P": {
			"Pedro": 1500,
			"Paulo": 1254,
		},
	}

	delete(funcionariosPorLetra, "P")

	for letra, funcionarios := range funcionariosPorLetra {
		fmt.Printf("Funcionários com a letra %s\n", letra)
		for nome, salario := range funcionarios {
			fmt.Printf("%s (Salário: R$%.2f)\n", nome, salario)
		}
		fmt.Println()
	}
}
