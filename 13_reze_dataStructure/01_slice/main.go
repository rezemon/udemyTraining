package main

import (
	"fmt"
)

func main() {

	var aSlice = []int{1, 3, 5, 7, 11, 18}
	fmt.Println("value of aSlice = ", aSlice)
	fmt.Print("type of aSlice =")
	fmt.Printf("%T", aSlice)
	fmt.Println()
	fmt.Print("length of aSlice =")
	fmt.Println(len(aSlice))
}
