package main

import (
	"fmt"
)

func main() {
	var n1 int
	n1 = 13
	n2 := factorial(n1)
	fmt.Println("factorial of ", n1, " = ", n2)
}
func factorial(aInt int) int {
	if aInt == 0 {
		return 1
	} else {
		return factorial(aInt-1) * aInt
	}
}
