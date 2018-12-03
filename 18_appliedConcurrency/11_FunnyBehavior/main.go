package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)

	}()

	//consume
	//reze : i keep making the mistake , put everything as goRoutine, like
	//the behavior vvv funny
	//for go run main.go  (will NOT NOT NOT see the Println)
	//for go run -race main.go (will see the Printlm)
	//next is to find out why ?
	go func() {
		for n := range c {
			fmt.Println(n)
		}

	}()

} //end main()
