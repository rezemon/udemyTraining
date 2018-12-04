package main

import (
	"errors"
	"log"

	"github.com/GoesToEleven/udemyTraining/_myTools/Describe"
)

//put Err infront
//it is idiomatic, so intellisense will help

var ErrNorgateMath = errors.New("norgath math: square root of negative number ")

func main() {
	myTools.Describe(ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(aNum float64) (float64, error) {
	if aNum < 0 {
		return 0, ErrNorgateMath
	}
	//implmentation
	return 42, nil
}

//see use of errors.New in the standard library
//http://golang.org/src/pkg/bufio/bufio.go
//http://golang.org/src/pkg/io/io.go
