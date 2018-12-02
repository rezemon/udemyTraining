package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println("some values in channel ", <-c)
}
