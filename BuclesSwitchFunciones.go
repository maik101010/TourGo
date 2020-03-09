package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

// func bucleFor(largoSuma int) int {
// 	sum := 0
// 	for i := 0; i < largoSuma; i++ {
// 		sum += i
// 	}
// 	return sum
// }

func bucleForDos() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func recursivo(valor int) int {
	if valor < 1 {
		return valor
	}
	fmt.Println(valor)
	return recursivo(valor - 1)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 9; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

//Ejemplo switch
func switcha() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func whereIsSaturday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchWithoutCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

//Vertices como arreglos
type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func principalBucles() {
	// fmt.Println(bucleFor(5))
	// bucleForDos()
	// recursivo(5)
	//--ejercicio poder
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )
	//Ejercicio for
	// fmt.Println(Sqrt(2))
	// fmt.Println("Sqrt: ", math.Sqrt(2))
	//Ejemplo Switch
	// switcha()
	// whereIsSaturday()
	// switchWithoutCondition()++
	//Ejemplo puntos
	// pointers()
	//Collections
	// v := Vertex{1, 2}
	// v.Y = 4
	// fmt.Println(v.Y)
	fmt.Println(v1, p, v2, v3)

}
