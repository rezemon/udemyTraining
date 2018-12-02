package main

import (
	"fmt"
	"sync"
)

var factCh chan int
var total int
var mutex sync.Mutex

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
			mutex.Lock()
			total = total * n
			fmt.Println("value of n = ", n)
			mutex.Unlock()
		}
	}()

	mutex.Lock()
	fmt.Println("factorial total = ", total)
	mutex.Unlock()
	//time.Sleep(2 * time.Second)
}
