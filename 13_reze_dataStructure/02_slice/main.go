package main

import (
	"fmt"
)

func main() {

	var aSlice = []string{"tan", "lee", "wong", "kim", "liang"}
	fmt.Println("value of aSlice =", aSlice)
	fmt.Println(aSlice[2]) //wong
	fmt.Println(aSlice[2:4])

	fmt.Println("bbAbb"[2]) //slice of the index position 2, which is 65 ascii for A

	fmt.Println("convert 65 to a string", string("bbAbb"[2])) //give back A

	fmt.Println("normal A")

}
