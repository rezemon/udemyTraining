package main

import "fmt"

func main() {

	var x = 0
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Printf("%T", increment)
	//above print confirm my thinking
	//the increment is of TYPE func() int
	//this way of declaring anonymous function is much much better
	/*
		this is the way programming should be
		the usage of ANONYMOUS FUNCTION
	*/
}
