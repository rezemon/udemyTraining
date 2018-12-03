//to compute 100 factorial concurrently
// create a list of channels
//almost like
// go factorial(x)
//go factorial(y)
//go factorial (z)

package main

import (
	"fmt"
)

func main() {
	c := createList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	factc := callFact(c)
	for n := range factc {
		fmt.Println("factorial of list = ", n)
	}
}

func createList(aList ...int) chan int {
	out := make(chan int)

	go func() {
		for _, value := range aList {
			out <- value
		}
		close(out)
	}()
	return out
}

func callFact(aCh chan int) chan int {
	out := make(chan int)
	go func() {

		for n := range aCh {
			out <- factorial(n)
		}
		close(out)
	}()
	return out
}

func factorial(aInt int) int {
	total := 1

	for i := 1; i <= aInt; i++ {
		total = total * i
	}
	return total
}
