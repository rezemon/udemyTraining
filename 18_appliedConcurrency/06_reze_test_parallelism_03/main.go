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

/*
my error , could be because the chan int, can put many things
in this scenario, i only want to have 1 thing
so, make(chan int,1)
*/
var aToken int
var count int
var wg sync.WaitGroup
var mutex sync.Mutex

var AB chan int
var BC chan int
var CA chan int

func main() {
	//after i limit it to a single buffer, it can run up to 647 times or 81 times, OCCASIONALLY
	AB = make(chan int, 1) //a channel
	BC = make(chan int, 1)
	CA = make(chan int, 1)
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
			mutex.Lock()
			count++
			mutex.Unlock()
		}
	}
}

func gorBput() {
	for n := range AB {
		if n == aToken {
			BC <- n
			fmt.Println("put to BC")
			mutex.Lock()
			count++
			mutex.Unlock()
		}
	}
}

func gorCput() {
	for n := range BC {
		if n == aToken {
			CA <- n
			fmt.Println("put to CA")
			mutex.Lock()
			count++
			mutex.Unlock()
		}
	}
}
