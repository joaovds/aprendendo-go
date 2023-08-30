package main

import "fmt"

type quadrado struct {
  lado float64
}

type circulo struct {
  raio float64
}

type figura interface {
  calcArea() float64
}

func main() {
  umQuadrado := quadrado{ 10 }
  umCirculo := circulo{ 5 }

  fmt.Println("Área do quadrado:", info(umQuadrado))
  fmt.Println("Área do círculo:", info(umCirculo))
}

func (q quadrado) calcArea() float64 {
  area := q.lado * q.lado
  return area
}

func (c circulo) calcArea() float64 {
  area := 3.14 * (c.raio * c.raio)
  return area
}

func info(f figura) float64 {
  return f.calcArea()
}

