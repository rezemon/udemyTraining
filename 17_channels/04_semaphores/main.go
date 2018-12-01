package main

import (
	"fmt"
)

// this dont work, so i try my idea in the next tutorial
//func init() {
//	fmt.Println("the first statement will execute here b4 main()")
//}
func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 11; i <= 21; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}
}
