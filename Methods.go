package main

import (
	"fmt"
	"math"
)

type MyVertex struct {
	X, Y float64
}

var entero = 1

func Abs(v MyVertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

// func (f MyFloat) Abs() float64 {
// 	if f < 0 {
// 		return float64(-f)
// 	}
// 	return float64(f)
// }

type MyInteger int64

func (i MyInteger) Abs() int64 {
	if i < 0 {
		return int64(-i)
	}
	return int64(i)
}

//Pointer receivers
func (v MyVertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//El asterisco mantiene el valor si es cambiado dentro de la referencia del metodo, sin el asterisco, el valor no se mantiene
func (v *MyVertex) Scale(f float64) {
	v.X = v.X * f //30
	v.Y = v.Y * f //40
}

//Otra forma de ser llamado
func Scale(v *MyVertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func mainMethods() {
	v := MyVertex{3, 4}
	//Formas de llamar metodo con Vertex
	// fmt.Println(v.Abs())
	// fmt.Println(Abs(v))
	//methods continued
	// f := MyFloat(-math.Sqrt2)
	// fmt.Println(f.Abs())
	// i := MyInteger(22)
	// fmt.Println(i.Abs)
	//--Operador asterisco *
	// v.Scale(10)
	// fmt.Println(v.Abs())
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
