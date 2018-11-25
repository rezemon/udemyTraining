package main

import (
	"fmt"
)

func main() {

	mySlice := make([]int, 2)
	fmt.Println(mySlice)

	fmt.Println(mySlice[1])

	mySlice[1] = 7
	fmt.Println(mySlice[1])

	mySlice[1]++

	fmt.Println(mySlice[1])

	fmt.Println(mySlice)

	mySlice = append(mySlice, 999)
	fmt.Println(mySlice)
}
