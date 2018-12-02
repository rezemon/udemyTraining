package main

import (
	"fmt"
)

func main() {
	//incrementor return a channel to whoever that call it
	incFoo := incrementor("Foo")
	incBar := incrementor("Bar")
	//puller take a channel as param and return a channel
	aResultFoo := puller(incFoo)
	aResultBar := puller(incBar)

	fmt.Println("result = ", <-aResultBar+<-aResultFoo)
}

// return a channel of integer
func incrementor(s string) chan int {
	incChan := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
			incChan <- 1
			fmt.Println(s, "index = ", i)
		}
		close(incChan)

	}()
	return incChan
}

func puller(aChan chan int) chan int {
	pullChan := make(chan int)
	go func() {
		var sum int
		for n := range aChan {
			sum = sum + n
		}
		//put into the channlel the sum
		pullChan <- sum
		close(pullChan)
	}()
	return pullChan
}
