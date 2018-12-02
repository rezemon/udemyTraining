package main

import "fmt"

func main() {
	//incrementor , create and return a channel
	aIncChan := incrementor()

	//puller, pass in chan as arguement, and return a channel from it
	aPullChan := puller(aIncChan)

	for n := range aPullChan {
		fmt.Println(n)
	}

}
func incrementor() <-chan int {
	incChan := make(chan int)
	go func() {
		for i := 0; i <= 10; i++ {
			incChan <- i
		}
		close(incChan) //close means still can read from it, but cannot write to it
	}()
	return incChan
}

func puller(aChan <-chan int) chan int {
	pullChan := make(chan int)
	go func() {
		for n := range aChan {
			pullChan <- n
		}
		close(pullChan) //close means still can read from it, but cannot write to it
	}()
	return pullChan
}

//https://golang.org/ref/spec#Channel_types
//https://golang.org/ref/spec#Channel_types
