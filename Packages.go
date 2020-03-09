package main

import (
	"fmt"
	"math"
	"math/rand"
)

const Pi = 3.14

func add(x int, y int) int {
	return x + y
}

func saludar(name string) string {
	return "Hola " + name
}

func moreOne(x, y, z int) (int, int, int) {
	x++
	y++
	z++
	return x, y, z
}

func split(sum int) (x, y, z int) {
	x = sum * 4 / 9
	y = sum - x
	z = 1
	return
}

func bucleFor(largoSuma int) int {
	sum := 0
	for i := 0; i < largoSuma; i++ {
		sum += i
	}
	return sum
}

// Para usar la palabra reservada var, se debe tener la variable afuera de una función, si no se usa, entonces nos da error de sintaxis
// var c, python, java bool
var i, j int = 1, 2

func funcionPrincipal() {
	fmt.Println("Mi número favorito es: ", rand.Intn(100))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("Suma de dos numeros", add(2, 1))
	fmt.Println(saludar("Michael"))
	//funcion que retorna multiples resultados
	x, _, _ := moreOne(1, 2, 3)
	fmt.Println(x)
	fmt.Println(split(17))
	k := 3
	i = 5
	//Con el var podemos asignar valores de dos formar "= := " pero sin el var, no, solamente ":="
	c, python, java := true, false, "no!"
	// var c, python, java = true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
	fmt.Println(bucleFor(5))
	// fmt.Println(i, j, c, python, java)
}
