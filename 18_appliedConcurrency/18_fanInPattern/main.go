package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("joe"), boring("ann"))
	//actually this is infinite, if i increase counter
	//it run forever,
	// reason that it stop, is because main() terminate
	// if increase it to infinite, it will be infinite
	// the loop is the one blocking the program from reaching main()
	for i := 0; i <= 38; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("you are both boring, i am leaving")
}

//this is NOT fan out
/*
fanout is
from a single channel, we the the input
and let it be processed by many func
e.g. would be a chan int (a channel of int)
and each integer, we sent it to a func factorial
*/
func boring(aStr string) chan string {
	out := make(chan string)

	go func() {
		//this is infinite
		for i := 0; ; i++ {
			out <- fmt.Sprintf("%s %d", aStr, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		//close(out)// unreacheable
	}()
	return out
}

//fan in
//multiple channel coming in
//output to a single channel
func fanIn(aCh1 chan string, aCh2 chan string) chan string {
	outCh := make(chan string)

	go func() {
		for {
			outCh <- <-aCh1
		}
		//close(outCh) //unreacheable
		//return outCh
	}()

	go func() {
		for {
			outCh <- <-aCh2
		}
		// close(outCh) //unreacheable
		//return outCh
	}()

	return outCh
}
