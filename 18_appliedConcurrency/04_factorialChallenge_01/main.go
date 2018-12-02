package main

import (
	"fmt"
)

func main() {
	var n1 float64
	n1 = 63
	n2 := factorial(n1)
	fmt.Println("factorial of ", n1, " = ", n2)
}
func factorial(aInt float64) float64 {
	if aInt == 0 {
		return 1
	} else {
		return factorial(aInt-1) * aInt
	}
}
