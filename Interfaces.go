package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t *T) M() {
	fmt.Println(t.S)
	t.S = "chao"
}

func M() {
	fmt.Println("segunda implementación de interfaz")
}

func mainInterfaces() {
	t := &T{"hello"}
	var i I = t
	i.M()
	fmt.Println(t.S)
	M()
	M()
}
