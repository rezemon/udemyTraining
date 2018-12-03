package main

import (
	"fmt"
)

func main() {

	c := gen(2, 3, 4, 5)
	cSq := sq(c)

	cSqAgain := sq(cSq)
	for n := range cSqAgain {
		fmt.Println("square of list = ", n)
	}
}

func gen(alist ...int) chan int {
	out := make(chan int)
	go func() {
		for _, value := range alist {

			out <- value
		}
		close(out)
	}()
	return out
}

func sq(aCh chan int) chan int {
	outSq := make(chan int)
	go func() {
		for n := range aCh {
			outSq <- n * n //square of n
		}
		close(outSq)
	}()
	return outSq
}
