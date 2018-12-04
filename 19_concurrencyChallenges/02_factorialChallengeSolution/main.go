//to compute 100 factorial concurrently
// create a list of channels
//almost like
// go factorial(x)
//go factorial(y)
//go factorial (z)

//change to execute 10,000 factorial concurrently
//and in parallel
//using fan in and fan out

package main

import (
	"fmt"
	"sync"
)

func main() {
	preC := listA()
	//c := createList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	c := createList(preC)
	/*
		what the difference between calling callFact(c) one time
		and calling callFact(c) Nth time
		remember, we will be taking off the number from the list
		this in fact are like Workers
		Each individual worker able to take things off from the channel
		look at the Worker and Publisher example
	*/
	/*
	   each statement here is sequential
	   but execution is probably parallel

	*/
	/*
	   fan out
	   multiple func reading from the same channel until that chanel is closed
	   reze: i guess it is close implicitly
	   distribute work across mutltiple funcs (TEN goRoutines) that all read from channel "c" above
	*/
	factc1 := callFact(c)
	factc2 := callFact(c)
	factc3 := callFact(c)
	factc4 := callFact(c)
	factc5 := callFact(c)
	factc6 := callFact(c)
	factc7 := callFact(c)
	factc8 := callFact(c)
	factc9 := callFact(c)
	factc10 := callFact(c)

	//then we need a merger
	//a fan in, gather all result together
	// this not so good example,
	//because these result are not related

	/*
	   fan in
	   multiplex multiple channels onto a single channel
	   merge the channels from factc1 ... factc10 onto a single channel
	*/
	var y int
	for n := range merge(factc1, factc2, factc3, factc4, factc5, factc6, factc7, factc8, factc9, factc10) {
		y++
		fmt.Println(y, "\t", n)
	}
	/*
		for n := range factc {
			fmt.Println("factorial of list = ", n)
		}
	*/
}

func listA() chan int {
	out := make(chan int)

	go func() {
		for i := 1; i <= 100000; i++ {
			for j := 20; j <= 30; j++ {
				//	for k := 3; k <= 12; k++ {
				out <- j
				//	}
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

//we are getting slice of (channel of int)
// so IMPORTANT IMPORTANT IMPORTANT
// can use for index,value range over index to get "channel of int"
func merge(manyCh ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	//all channels have to complete
	wg.Add(len(manyCh))

	//we range over channel and put them into a SINGLE channel
	//we dont need index
	//value will hold the individual channel
	for _, value := range manyCh {
		go func(aCh chan int) {
			//out <- <-aCh // the int from aCh, will be pass in to out channel
			for n := range aCh {
				out <- n
			}
			wg.Done()
		}(value)

	}
	//but to merge, have to wait for all channels to complete putting their value to out
	//so need a snyc.WaitGroup

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
