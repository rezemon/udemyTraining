package main

import (
	"fmt"
)

func main() {

	//try out if it is a slice consisting of slice of string []string
	//==> [] []string
	//or it is my idea
	// [][]string, a 2 dimension of string

	/*

	   tutorial could be right, it is a slice, contaning a slice
	   because we cannot put the size there,
	   i will try to see if can do it here
	*/
	//var aMultiDimenSlice = make([][]int, 3)
	var aMultiDimenSlice = make([][]int, 9) // in a slice, put 9 slice

	fmt.Printf("%T", aMultiDimenSlice)
	fmt.Println()
	fmt.Println(aMultiDimenSlice)
	//aMultiDimenSlice[1][1] = 5 //ERORR

	fmt.Println(aMultiDimenSlice)

}
