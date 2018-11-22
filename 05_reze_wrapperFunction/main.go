/*
this is the first time that i understand WRAPPER FUNCTION
it is something that WRAP around AND return the FUNCTION
*/

package main

import "fmt"

/*
i am declaring a function call myWrapper()
the return type of type func() int

and AGAIN, we make use of ANONYMOUS func() to implement the WRAPPER concept

*/
func myWrapper() func() int {
	var x = 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := myWrapper()
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Printf("%T\n", myWrapper)
	//above should say that myWrapper is a function that is of type int
	// func() int XXXXXX WRONG
	//myWrapper is a func that return type func() int
	//hence it import
	//"func()" "func() int"
	//above, it is a "func()"
	// that is returning "func() int"

	fmt.Printf("%T\n", increment)
	//above should say that myWrapper is a function that is of type int
	// func() int

	/*

	 */
}
