package main

import (
	"fmt"
	"sync"
	"time"
)

/*
3 go routine execute
the main() goroutine
the foo() goroutine
the bar() goroutine
we cannot see the println, we add waitgroup
*/

var wg sync.WaitGroup

func main() {
	fmt.Println("concurrency")
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait()
}
func foo() {
	for i := 0; i <= 30; i++ {
		fmt.Println(i, "foo")
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i <= 30; i++ {
		fmt.Println(i, "bar")
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}
