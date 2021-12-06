package formas

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

func EscreverArea(f Forma) {
	fmt.Printf("A área da forma é %0.2f\n", f.Area())
}

type Retangulo struct {
	Altura  float64
	Largura float64
}

// Aqui conectei a struct retangulo a interface forma ao implementar o "metodo" area para retangulo
func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

type Circulo struct {
	Raio float64
}

// Aqui conectei a struct circulo a interface forma ao implementar o "metodo" area para circulo
func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio, 2)
}
