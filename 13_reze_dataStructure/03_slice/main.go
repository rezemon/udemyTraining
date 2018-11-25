package main

import "fmt"

func main() {
	//make a slice of size 50, with underlying array size =100
	//every slice has a underlying array
	//something that the slice is pointing to where the data is stored
	//hence we call slice a REFERENCE type
	//map is a REFERENCE type
	// channel is also a REFERENCE type
	/*
	  the reason that we use make is probably because
	  you have to make it
	  some setup that computer has to do, to make slice work
	*/
	/*
	   slice is built on top of an underlying array
	   it is pointing to that array
	*/
	var aSliceA = make([]int, 20, 100)

	//same as aSliceA, we new a array of size 100, and slice the first 50 position
	var aSliceB = new([100]int)[0:20]

	aSliceA[7] = 7
	aSliceB[9] = 9
	//above when printed should nicely show the positioning
	fmt.Println("aSliceA =", aSliceA)
	fmt.Println("aSliceB =", aSliceB)
	//above VERY VERY WELL DEMOSTRATED, BETTER THAN THE TEACHER IN THE TUTORIAL
	//it show 20 items in the slice, though underlying array is of size 100
}
