package main

import (
	"fmt"
	"sync"
)

// this dont work, so i try my idea in the next tutorial
func init() {
	fmt.Println("the first statement will execute here b4 main()")
}
func main() {

	c := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		//wg.Add(1)
		for i := 0; i <= 10; i++ {
			c <- i
		}
		wg.Done()
	}()

	go func() {
		//wg.Add(1)
		//for i := 0; i <= 10; i++ {
		//	c <- i
		//}
		for i := 11; i <= 21; i++ {
			c <- i
		}

		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
