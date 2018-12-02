package main

import (
	"fmt"
)

var factCh chan int
var total int

func main() {
	//fmt.Println("start")
	factCh = make(chan int)

	total = 1

	go func() {
		for i := 1; i <= 5; i++ {
			factCh <- i // put 5 values inside
		}
		close(factCh)
	}()

	// retrieve it like a stack
	go func() {
		for n := range factCh {
			total = total * n
			fmt.Println("value of n = ", n)
		}
	}()

	fmt.Println("factorial total = ", total)
	//time.Sleep(2 * time.Second)
}
