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
	preC := listA()
	//c := createList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	c := createList(preC)
	factc := callFact(c)
	for n := range factc {
		fmt.Println("factorial of list = ", n)
	}
}

func listA() chan int {
	out := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			for j := 3; j <= 12; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func createList(aList chan int) chan int {
	out := make(chan int)

	go func() {
		for n := range aList {
			out <- n
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

//the factorial here can be recursive
/*
previously , when i thought about the factorial,
i was thinking of too fine details
here, parallel 100 execution of factorial
for factorial(13), we let golang concurrency take over
of course, if factorial(13) is very expensive
we can optimize it like quick sort
break into half, left sort, right sort

*/
func factorial(aInt int) int {
	total := 1

	for i := 1; i <= aInt; i++ {
		total = total * i
	}
	return total
}
