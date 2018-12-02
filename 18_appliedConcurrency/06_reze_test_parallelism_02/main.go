/*
if we try this sort of channel communication
on a unbuffered communication
we do a token passing
ONE and only ONE token passed among , A, B, C
A pass to B
B pass to C
C pass to A
keep repeating and see what happen

initial condition is to put a token
exit condition is to remove the token
*/

package main

import (
	"fmt"
	"sync"
)

var aToken int
var count int
var wg sync.WaitGroup
var mutex sync.Mutex

var AB chan int
var BC chan int
var CA chan int

func main() {
	AB = make(chan int) //a channel
	BC = make(chan int)
	CA = make(chan int)
	aToken = 8
	//aCh := make(chan int)

	//initial condition put a token inside
	go func() {
		AB <- aToken
	}()
	//fmt.Println("this is start")

	go gorAput()
	go gorBput()
	go gorCput()

	fmt.Println("value of token = ", aToken)
	fmt.Println("value of count =", count)
}

//what A receive, A pass to B
//what B receive, B pass to C
//what C receive, C pass to A
//dont need 3 channel, one channel is good enough
// wrong , 3 channel

// i can put, provided someone ready to get
// i can get, provided someone ready to put
//but async => means i just put
//but async => means i just get
//but async => means i just put , without concern whether someone ready to get
//but async => means i just get , without concern whether someone ready to put

//if there is something to read from channel CA, i will pass it to channel AB
//if there is something READ from channel, then i will pass it to channel AB, else keep reading

//experiment FAILED, i am not good with channel yet

func gorAput() {
	for n := range CA {
		if n == aToken {
			AB <- n
			fmt.Println("put to AB")

		}
	}
}

func gorBput() {
	for n := range AB {
		if n == aToken {
			BC <- n
			fmt.Println("put to BC")

		}
	}
}

func gorCput() {
	for n := range BC {
		if n == aToken {
			CA <- n
			fmt.Println("put to CA")
		}
	}
}
