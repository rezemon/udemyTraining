package main

import (
	"fmt"
	"sync"
)

func main() {

	in := gen()

	// FAN OUT
	// Multiple functions reading from the same channel until that channel is closed
	// Distribute work across multiple functions (ten goroutines) that all read from in.
	xc := fanOut(in, 10)

	// FAN IN
	// multiplex multiple channels onto a single channel
	// merge the channels from c0 through c9 onto a single channel

	//#########################
	//TROUBLESHOOTING
	//#######################
	/*
		TROUBLESHOOTING because from the len of xc, we are expecting len=10
		it show us len=20
		so something must be wrong,
		reze: this kinda suit me, the debugger cannot help much
		reze: it need clear thing
		that is how the speaker or teacher figure out
		cause he was seeing the size of the xc NOT what he expect
	*/
	fmt.Printf("%T\n", xc)
	fmt.Println("*********** length of xc", len(xc))
	for _, value := range xc {
		fmt.Println("***************", <-value)
	}

	for n := range merge(xc...) {
		fmt.Println(n)
	}

}

func gen() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func fanOut(in <-chan int, n int) []<-chan int {
	//################################################
	//xc := make([]<-chan int, n) //THIS IS THE ERROR
	//################################################
	// because, it is make of size n, which is 10
	//and append, will start from 11,12,13... appending to it
	//when program access it, it will access the first 10, which is nil, so there is error
	// how to debug ?
	// not easy as, the debugger is not so fine tune
	var xc []<-chan int //CORRECT

	for i := 0; i < n; i++ {
		xc = append(xc, factorial(in))
	}
	return xc
}

func factorial(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

/*
CHALLENGE #1:
-- This code throws an error: fatal error: all goroutines are asleep - deadlock!
-- fix this code!
*/
