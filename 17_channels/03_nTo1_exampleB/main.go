package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// this dont work, so i try my idea in the next tutorial
func init() {
	fmt.Println("the first statement will execute here b4 main()")
	wg.Wait()
}

func main() {

	c := make(chan int)

	go func() {
		wg.Add(1)
		for i := 0; i <= 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 0; i <= 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		//	wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
