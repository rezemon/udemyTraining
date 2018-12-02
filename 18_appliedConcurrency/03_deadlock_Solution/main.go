package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
		}
		close(c)
	}()

	//fmt.Println("channel c contain ", <-c)
	//for is very important for channel
	for n := range c {
		fmt.Println("channel c contain = ", n)
	}
	// in order to range over a channel, the channel has to be CLOSED
}
