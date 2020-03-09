package main

import (
	"fmt"
	"math"
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

func main() {
	// fmt.Println(bucleFor(5))
	// bucleForDos()
	// recursivo(5)
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
