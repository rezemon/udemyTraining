package main

import "fmt"

// below function return nothing
func zero(x int) {
	x = 99
}

// below is a function that accept an ADDRESS POINTER
func zeroB(x *int) {
	/*
		pass here is the pointer to var y
		thus we DE-REFERENCE it by adding * infront,
	*/
	*x = 88
}
func main() {
	var y = 33
	fmt.Println("original value of y = ", y)
	zero(y)
	fmt.Println("value of y after passed by value IS NOT AFFECTED = ", y)
	// y is passed by value to func zero, thus y remain unchanged

	zeroB(&y)
	fmt.Println("value of y after passed by pointer IS CHANGED = ", y)
}

// above show we can have very PERFORMANT func
// we dont pass the whole thing about in the memory
// we pass the pointer
// REMEMBER, pointer is pass by value
// pointer is just a value
