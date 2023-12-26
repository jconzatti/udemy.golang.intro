package area

import "math"

//Proporção entre a circunferência e o diâmetro de um circulo
const PI = 3.1416

//Área de um círculo com o raio
func Circulo(raio float64) float64 {
	return math.Pow(raio, 2) * PI
}

//Área de um retângulo com base e altura
func Retangulo(base, altura float64) float64 {
	return base * altura
}

//Área de um triângulo com base e altura
func _Triangulo(base, altura float64) float64 {
	return base * altura / 2
}
