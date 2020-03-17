package main

import (
	"fmt"
	"math"
)

type VertexPointer struct {
	X, Y float64
}

//Funcion que recibe puntero y float para operar primera forma
func (v *VertexPointer) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//Funcion que recibe puntero y float para operar segunda forma
func ScaleFunc(v *VertexPointer, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v VertexPointer) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v VertexPointer) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func mainMethodsPointer() {
	//Dos formas de enviar el puntero a metodo receptor
	//Primera forma
	// v := VertexPointer{3, 4}
	// v.Scale(2)        //v.x = 6, v.y = 8
	// ScaleFunc(&v, 10) //v.x = 60, v.y=80
	// //Segunda forma
	// // p := &v
	// p := &VertexPointer{4, 3}
	// p.Scale(3)      //p.x = 12, p.y = 9
	// ScaleFunc(p, 8) //p.x = 96, p.y=72
	// fmt.Println(v, p)

	//Methos and pointer indirection(2)
	v := VertexPointer{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &VertexPointer{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))

}
