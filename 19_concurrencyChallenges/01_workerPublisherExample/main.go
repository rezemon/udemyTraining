package main

import (
	"fmt"
	"time"
)

var workerID int
var publisherID int

func main() {
	input := make(chan string)
	go workProcess(input)
	go workProcess(input)
	go workProcess(input)

	go publisher(input)
	go publisher(input)
	go publisher(input)
	go publisher(input)

	time.Sleep(10 * time.Second)
}

func publisher(out chan string) {
	publisherID++
	thisID := publisherID
	dataID := 0
	for {
		dataID++
		fmt.Printf("publisher %d is pushing data\n", thisID)
		data := fmt.Sprintf("data from publisher %d. Data %d", thisID, dataID)
		out <- data
	}
}

//this is definitely fan out
func workProcess(in chan string) {
	workerID++
	thisID := workerID
	for {
		fmt.Println("waiting for input = ", thisID)
		input := <-in
		fmt.Println(thisID, " input is = ", input)
	}
}
