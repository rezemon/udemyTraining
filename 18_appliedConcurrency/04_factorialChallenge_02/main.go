package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	var n1 int

	wg.Add(1)

	n1 = 6
	n2 := factorial(n1)
	fmt.Println("factorial of ", n1, " = ", n2)

	//using go rountine can only achieve concurrency
	//using channels, we achieve parallelism
	//factorial has no meaning in concurrency as result need to wait for every piece
	//but if parallelism, yes it will speed up
	/*
	   f(5) = f(4) * 5
	   f(4) = f(3) * 4
	   f(3) = f(2) * 3
	   f(2) = f(1) * 2
	   f(1) = f(0) * 1
	   f(0) = 1
	   if each of the factorial is handled in PARALLEL, it speed up the calculation
	*/
	/*
	   a channel is how goRoutine talk to each other to achieve parallelism
	   we spawn a f(5), wait for the result
	   f(5), spawn f(4), wait for the result,
	   f(4), spawn f(3), wait for the result, so on and so forth
	*/

	go test()
	wg.Wait()
}

/* // cannot be put here
go func test(){
	fmt.Println("this is a test")
}()
*/

func test() {
	fmt.Println("this is a test")
	wg.Done()
}

func factorial(aInt int) int {
	if aInt == 0 {
		return 1
	} else {
		return factorial(aInt-1) * aInt
	}
}
