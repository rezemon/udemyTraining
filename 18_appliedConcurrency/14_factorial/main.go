package main

import (
	"fmt"
)

func main() {
	// factorial return a channel
	//this show multiple way of using channel
	c := factorial(4)

	for n := range c {
		fmt.Println("factorial of 4 = ", n)
	}
}
func factorial(aInt int) chan int {
	out := make(chan int)
	go func() {
		total := 1
		for i := 1; i <= aInt; i++ {
			total = total * i
		}
		out <- total
		close(out)
	}()
	return out
}
