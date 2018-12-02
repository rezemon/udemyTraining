package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	/*
			cannot put something inside channel
		  unless there is somebody WAITING to remove something from it
			so solution is to make it concurrent using keyword "go func"
	*/
	//go func() {
	c <- 1
	//}()
	fmt.Println("some values in channel ", <-c)
}
