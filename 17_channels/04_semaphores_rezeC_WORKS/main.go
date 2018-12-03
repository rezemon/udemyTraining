package main

import (
	"fmt"
	"sync"
)

// this dont work, so i try my idea in the next tutorial
//func init() {
//	fmt.Println("the first statement will execute here b4 main()")
//}

var mutex sync.Mutex

func main() {

	c := make(chan int)

	cHelper := 0
	//because cHelper is shared
	// and two go routine is accessing it, thus race condition
	//so we make have to have TWO cHelper , or something else

	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
		}
		mutex.Lock()
		cHelper++
		mutex.Unlock()
	}()

	go func() {
		for i := 11; i <= 21; i++ {
			c <- i
		}
		mutex.Lock()
		cHelper++ // after adding mutex , found only 1 race data
		mutex.Unlock()
	}()

	go func() {
		for {
			mutex.Lock()
			if cHelper == 2 {

				fmt.Println("closed how many times") //close 2 times
				close(c)
				//cHelper = 88
				break //to make sure close(c) execute 1 time only
			}
			mutex.Unlock() //no more error, but this show, concurrent running goRoutine, as long as access a SHARED resource, will have race conditions
		}
	}()

	for n := range c {
		fmt.Println(n)
	}
	fmt.Println("cHelper = ", cHelper)
}
