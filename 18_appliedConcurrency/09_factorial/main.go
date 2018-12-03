package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	//done := make(chan bool)
	go func() {
		for i := 1; i <= 6; i++ {
			c <- i
		}
		close(c)
		//done <- true
	}()

	total := 1
	for n := range c {
		total = total * n
		fmt.Println("n =", n)
	}

	fmt.Println("total = ", total)
	//done <- true
}
