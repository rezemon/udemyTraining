package main

import (
	"fmt"
	"math/rand"
	"runtime"
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
var counter int

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	fmt.Println("concurrency")
	wg.Add(2)
	go incrementor("Foo : ")
	go incrementor("Bar : ")
	wg.Wait()
}
func incrementor(s string) {
	for i := 0; i <= 20; i++ {
		x := counter
		x++
		//time.Sleep(time.Duration(3 * time.Millisecond))
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter :", counter)
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
