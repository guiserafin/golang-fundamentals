package formas

import (
	"fmt"
	"math"
)

// interface so possui assinaturas de metodos
type Forma interface {
	Area() float64
}

func EscreverArea(f Forma) {
	fmt.Printf("A área da forma é %0.2f \n", f.Area())
}

type Retangulo struct {
	altura  float64
	largura float64
}

func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

type Circulo struct {
	raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * (c.raio * c.raio)
}
