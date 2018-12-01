package main

import (
	"fmt"
)

func main() {

	c := make(chan int) //unbuffer

	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
		}
		close(c) //after close no one can put anything on channel c
		// but still can receive from channel c
	}()

/*
RANGE clause is vvvvvvvvv important for dealing with goRoutine

 */
	for n := range c { // n is receiving from c, it BLOCK or stop if nothing is send
		fmt.Println(n) // it do this until the channel is close, and nothing left on channel
	}

	// time.Sleep(time.Second) //we dont need this becoz see above

}
