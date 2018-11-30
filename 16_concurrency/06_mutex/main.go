//mutex
//mutually exclusive

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
var mutex sync.Mutex

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	fmt.Println("concurrency")
	wg.Add(2)
	go incrementor("Foo : ")
	go incrementor("Bar : ")
	wg.Wait()
	fmt.Println("final counter =", counter)
}
func incrementor(s string) {
	for i := 0; i <= 20; i++ {
		//time.Sleep(time.Duration(3 * time.Millisecond))
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		mutex.Lock()
		//x := counter
		//x++
		//counter = x
		counter++
		fmt.Println(s, i, "Counter :", counter)
		mutex.Unlock()

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
