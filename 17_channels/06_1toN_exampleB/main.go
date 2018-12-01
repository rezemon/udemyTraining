package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	c := make(chan int)
	done := make(chan bool)
	n := 40

	//putting things to channel
	go func() {
		for i := 0; i <= 100000; i++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i <= n; i++ {
		go func() {
			for n := range c {
				fmt.Println(n)
			}
			done <- true
		}()
	}

	for i := 0; i <= n; i++ {
		<-done
	}
	elapsed := time.Since(start)
	log.Printf("it took %s", elapsed)
}
