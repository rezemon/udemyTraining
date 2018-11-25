package main

import (
	"fmt"
)

func main() {

	var aArray [256]int

	for i := 0; i <= 255; i++ {
		aArray[i] = i
	}
	fmt.Print("aArray = ", aArray)

	for _, v := range aArray {
		fmt.Printf("%v - %T - %b", v, v, v)
		fmt.Println()
	}

}
