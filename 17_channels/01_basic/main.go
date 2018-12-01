package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int) //unbuffer

	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
		}
	}()

	go func() {
		for i := 0; i <= 10; i++ {
			fmt.Println(<-c)
		}
	}()

	time.Sleep(time.Second)

}
