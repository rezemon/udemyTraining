package main

import "fmt"

type I interface {
	M() //method
}

type T struct {
	S string
}

type TT struct {
	I int
}

// related to struct T, implemented method M()
//therefore T is interface I
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func (t *TT) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.I)
}

func main() {
	var i I //i is interface I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	i=&TT{1}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
