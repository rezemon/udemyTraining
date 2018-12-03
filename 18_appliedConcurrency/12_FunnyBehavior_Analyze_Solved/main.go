package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
		done <- true
	}()

	//consume
	//reze : i keep making the mistake , put everything as goRoutine, like
	//the behavior vvv funny
	//for go run main.go  (will NOT NOT NOT see the Println)
	//for go run -race main.go (will see the Printlm)
	//next is to find out why ?
	//because, goRoutine, execute it and it is already completed
	go func() {
		for n := range c {
			fmt.Println(n)
		}
		done <- true
	}()

	//this is working, the moment, i make the <-done as go routine, problem problem appear again

	<-done
	<-done
} //end main()
