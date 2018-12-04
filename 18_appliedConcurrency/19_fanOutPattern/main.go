package main

import (
	"fmt"
	"sync"
)

func main() {

	//c := gen(2, 3, 4, 5)
	c := gen(5, 4, 3, 2)

	//fanout
	//distribute the sq work across two go routine that both read from channel c above
	//reze: however my view is that it is still sequential,
	//if we go func() factorial , we loop over the list and send,
	//then all is concurrent
	//BUT BUT BUT interesting, there is only 1 channel c
	// if we take things off it,
	// it will be random
	// cSq may take 2,5
	// cSq may take 3,4
	//OR
	// cSq may take 2,4,5
	// cSq may take 3
	//OR
	// cSq may take 0
	// cSq may take 2,3,4,5
	//the func pull the info from the channel, whoever func is ready, it will pull
	//ANY combination could be possible ???

	//MULTIPLE func pulling from 1 channel, this is fanOut
	//this is definitely fanOut
	cSq := sq(c)
	cSq2 := sq(c) //it may not be the same function , one could be sqRoot
	//################################
	//################################
	// but remember , a channel is a SEQUENTIAL SERIAL pipeline
	//################################
	//################################

	/*
	   when i range over a channel of int
	   i get integer

	   when i use the
	   for index, value :=  range channel of int
	   the value contain "channel of int"
	*/
	// now to do fanIn
	for n := range merge(cSq, cSq2) {
		fmt.Println(n)
	}

	/* from old tutorial, obsolete
	for n := range cSq {
		fmt.Println("square of list = ", n)
	}
	*/
}

func gen(alist ...int) chan int {
	out := make(chan int)
	go func() {
		// here is ranging over a slice, that is why have index, value
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

//many channel, ONE channel out => fan in
// the param is a slice
func merge(manyChanIn ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(manyChanIn))

	/*
	   BELOW VERY VERY IMPORTANT
	   CLOSURE
	   if we dont go func(PARAM){}(ARGU for PARAM)
	   then we are ALWAYS refering to the SAME aCh
	   for all the goRoutine that is launch
	   listen to the video again if dont understand
	*/
	//for oneCh := range manyChanIn {
	for _, value := range manyChanIn {
		go func(aCh chan int) {
			for n := range aCh {
				out <- n
			}
			wg.Done()
		}(value) //here value is a (chan of int) , probably because i send a slice of "chan of int"
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
