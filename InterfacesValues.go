package main

import (
	"fmt"
	"math"
)

type Interface interface {
	M()
}

type Struct struct {
	S string
}

// func (t *Struct) M() {
// 	fmt.Println(t.S)
// }

func (t *Struct) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i Interface) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeInterfaceEmpty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func mainInterfacesValues() {
	var i Interface
	var t *Struct
	i = t
	//Toma el metodo por herencia, teniendo en cuenta el parametro enviado.
	// i = &Struct{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	//empty interface
	fmt.Println("--------Empty interface------")
	var iEmpty interface{}
	describeInterfaceEmpty(iEmpty)

	iEmpty = 42
	describeInterfaceEmpty(iEmpty)

	//Type assertions
	fmt.Println("--------Type assertions------")
	var iAssert interface{} = "hello"
	s := iAssert.(string)
	fmt.Println(s)

	s, ok := iAssert.(string)
	fmt.Println(s, ok)

	f, ok := iAssert.(float64)
	fmt.Println(f, ok)

	// f = iAssert.(float64) // panic
	// fmt.Println(f)
	//Type interface
	fmt.Println("--------//Type interface------")
	do(21)
	do("hello")
	do(true)

}
