package main

import "math"

//Ponto representa uma coordenada no plano cartesiano
type Ponto struct {
	x float64
	y float64
}

func obterCatetos(a, b Ponto) (catetoX, catetoY float64) {
	catetoX = math.Abs(b.x - a.x)
	catetoY = math.Abs(b.y - a.y)
	return
}

//Dinstância entre dois pontos no plano cartesiano
func Distancia(a, b Ponto) float64 {
	cx, cy := obterCatetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
