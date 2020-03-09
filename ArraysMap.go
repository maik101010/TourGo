package main

import (
	"fmt"
	"strings"
)

func arrayInt() [6]int {
	// var a [2]string
	// a[0] = "Hello"
	// a[1] = "World"
	// fmt.Println(a[0], a[1])
	// fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// var s []int = primes[2:3]
	return primes
}

func arrayString() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

func arrayBoolean() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

// type Vertex struct {
// 	Lat, Long float64
// }

// var m map[string]Vertex
//Función que cuenta la cantidad de palabras repetidas
func WordCount(s string) map[string]int {
	var arregloString = strings.Fields(s)
	mapa := make(map[string]int)
	for i := 0; i < len(arregloString); i++ {
		v := mapa[arregloString[i]]
		if v != 0 {
			mapa[arregloString[i]]++
		} else {
			mapa[arregloString[i]] = 1
		}
	}
	return mapa
}

func principal() {
	// var mapa = WordCount("hola mundo")
	// // m = make(map[string]Vertex)
	// var m = map[string]Vertex{
	// 	"Bell Labs": {40.68433, -74.39967},
	// 	"Google":    {37.42202, -122.08408},
	// }
	// m["asaaaa Labs"] = Vertex{
	// 	40.68433, -74.39967,
	// }
	// m["Venadillo"] = Vertex{
	// 	72.1212, -44.12222,
	// }
	// fmt.Println(m["Google"])
	fmt.Println(WordCount("A man a plan a canal panama."))
	//Asignando valores a arreglo por defecto y recorrerlo
	// pow := make([]int, 12) //Hacemos un arraglo de 10 enteros 0
	// for i := range pow {
	// 	pow[i] = 1 << uint(i) // == 2**i
	// 	fmt.Printf("%d\n", pow[i])
	// }
}
