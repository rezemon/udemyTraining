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
	"sync"
	"time"
)

var mutex sync.Mutex
var aToken1 int

//var aToken2 int
var Ch1 chan int
var Ch2 chan int

//to stop everything gracefully, so to speak
var ChStop chan int
var ChStopStr string // init to a value that will change later to cause stop
var count int

func main() {
	Ch1 = make(chan int)
	Ch2 = make(chan int)
	ChStopStr = "continue"
	aToken1 = 88
	//aToken2 = 99
	go func() {
		Ch1 <- aToken1
		//Ch2 <- aToken2
	}()

	go gorAput()
	go gorBput()

	/*
	   we have to wait this timing then can see the running
	   so how to use channel to ensure everything is sync
	   at some point, another channel to ensure everything is completed

	*/
	//	time.Sleep(100 * time.Second) //100 sec
	//time.Sleep(3 * time.Second) //10 sec
	//mutex.Lock()

	/*
		go func() {

			//<-ChStop // after some count, we put something into CChStop, so that this will step over when it is read
			//fmt.Println("count = ", count)
			for n := range ChStop {
				fmt.Println("value of n = ", n)
			}
		}()
	*/

	//BELOW NOT GOOD, STILL WORK IN PROGRRESS
	go func() {
		n := <-ChStop
		fmt.Println(n)
	}()
	//hmm... comment off, wont see anything ???
	//fmt.Println("count = ", count) //here still a race condition // cannot be here, it will not be seen
	//mutex.Unlock()

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
		mutex.Lock()
		count++
		if count >= 888 && ChStopStr == "continue" {
			//ChStop <- count
			ChStop <- 77       //any number
			ChStopStr = "Stop" //this enter once only
		}
		mutex.Unlock()
		fmt.Println("put to Ch1")
	}
	//go gorAput() //this put or dont put same, it will run
	//the exit is due to the exit or the main program, this piece of code works!!!
	//if we pass back and forth correctly, there should not be a Race condition
	// we will verify this with a shared resources, e.g. a global field
	//lee chuan boon , you are solid, there is no race condition
	//WRONG, there is when i try to print count
}

func gorBput() {
	// if A give me something, i pass it back to A via Ch2
	//if ch2 has something, do nothing
	//if ch2 has nothing, put something
	for n := range Ch1 {
		Ch2 <- n
		mutex.Lock()
		count++
		mutex.Unlock()
		fmt.Println("put to Ch2")
	}
	//go gorBput()

}

func gorAget() {
}

func gorBget() {
}
