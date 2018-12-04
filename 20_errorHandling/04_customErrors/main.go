package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(aNum float64) (float64, error) {
	if aNum < 0 {
		return 0, errors.New("norgath math: square root of negative number ")
	}
	return 42, nil
}

//see use of errors.New in the standard library
//http://golang.org/src/pkg/bufio/bufio.go
//http://golang.org/src/pkg/io/io.go
