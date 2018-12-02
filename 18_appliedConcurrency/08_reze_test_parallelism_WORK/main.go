//try easier
//one channel
// A put, B read
// B read, will put back for A to read
// causing a infinite loop

//cannot 1 channel
//2 channel
//channel #1, from A to B
//chanell #2, from B to A

//####################
//WORKS WORKS WORKS
//####################

package main

import (
	"fmt"
	"time"
)

var aToken1 int
var aToken2 int
var Ch1 chan int
var Ch2 chan int

func main() {
	Ch1 = make(chan int)
	Ch2 = make(chan int)
	aToken1 = 88
	aToken2 = 99
	go func() {
		Ch1 <- aToken1
		//Ch2 <- aToken2
	}()

	go gorAput()
	go gorBput()

	//	time.Sleep(100 * time.Second) //100 sec
	time.Sleep(10 * time.Second) //10 sec
}

//goR == goRoutine
//if something on channel, read it
func gorAput() {
	// if B give me something, i pass it back to B via Ch1
	//if ch1 has something, do nothing
	//if ch1 has nothing, put something

	// i range over Ch2, if found something, put in Ch1
	for n := range Ch2 {
		Ch1 <- n
		fmt.Println("put to Ch1")
	}
	//go gorAput() //this put or dont put same, it will run
	//the exit is due to the exit or the main program, this piece of code works!!!
}

func gorBput() {
	// if A give me something, i pass it back to A via Ch2
	//if ch2 has something, do nothing
	//if ch2 has nothing, put something
	for n := range Ch1 {
		Ch2 <- n
		fmt.Println("put to Ch2")
	}
	//go gorBput()

}

func gorAget() {
}

func gorBget() {
}
