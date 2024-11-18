package main

import (
	"math"
)

// Triangulo es el objeto (se ponen las propiedades, es un conjunto de propiedades).
type Triangulo struct {
	Lado1 float64
	Lado2 float64
	Lado3 float64
}

// GetTipo para saber el tipo de triangulo.
func (triangulo *Triangulo) GetTipo() string {
	tipo := "Desconocido"

	if triangulo.Lado1 == triangulo.Lado2 && triangulo.Lado2 == triangulo.Lado3 {
		tipo = "Equilátero"
	} else if triangulo.Lado1 != triangulo.Lado2 && triangulo.Lado1 != triangulo.Lado3 && triangulo.Lado2 != triangulo.Lado3 {
		tipo = "Escaleno"
	} else {
		tipo = "Isósceles"
	}

	if triangulo.EsRectangulo() {
		tipo += " Rectángulo"
	}

	return tipo
}

// EsRectangulo ver si es rectangulo.
func (triangulo *Triangulo) EsRectangulo() bool {
	return (math.Pow(triangulo.Lado1, 2) == (math.Pow(triangulo.Lado2, 2) + math.Pow(triangulo.Lado3, 2))) ||
		(math.Pow(triangulo.Lado2, 2) == (math.Pow(triangulo.Lado1, 2) + math.Pow(triangulo.Lado3, 2))) ||
		(math.Pow(triangulo.Lado3, 2) == (math.Pow(triangulo.Lado1, 2) + math.Pow(triangulo.Lado2, 2)))
}

// SetLado1 mete lado1.
func (triangulo *Triangulo) SetLado1(lado1 float64) {
	triangulo.Lado1 = lado1
}

// SetLado2 mete lado2.
func (triangulo *Triangulo) SetLado2(lado2 float64) {
	triangulo.Lado2 = lado2
}

// SetLado3 mete lado3.
func (triangulo *Triangulo) SetLado3(lado3 float64) {
	triangulo.Lado3 = lado3
}

// GetLado1 devuelve lado1.
func (triangulo *Triangulo) GetLado1() float64 {
	return triangulo.Lado1
}

// GetLado2 devuelve lado2.
func (triangulo *Triangulo) GetLado2() float64 {
	return triangulo.Lado2
}

// GetLado3 devuelve lado3.
func (triangulo *Triangulo) GetLado3() float64 {
	return triangulo.Lado3
}
