package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	/* //this cause wg.Done() extra by 1
	var n1 int
	n1 = 13
	n2 := factorial(n1)
	fmt.Println("factorial of ", n1, " = ", n2)
	*/

	/*
		   to show concurrency
		   factorial(13) is the longest time taken
		   factorial(9) 2nd longest time taken
		   factorial(3) is the shortest
		   thus we expect factorial(3) to finish executing, if there is concurrency

			 actual run, the completion time is random, sometimes, factorial(13) will complete first
			 sometimes, factorial(3) complete first
			 if change to factorial(63)??
			 still it is random
	*/
	wg.Add(3)
	go factorial(63, 63)
	go factorial(9, 9)
	go factorial(3, 3)
	wg.Wait()
}
func factorial(aInt int, aIntFix int) int {
	if aInt == 0 {
		fmt.Println("exit case for factorial = ", aIntFix)
		wg.Done()
		return 1
	} else {
		return factorial(aInt-1, aIntFix) * aInt
	}
}
