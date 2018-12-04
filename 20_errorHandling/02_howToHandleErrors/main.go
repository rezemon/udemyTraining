package main

import (
	"fmt"
	"log"
	"os"
)

//we can set where standard error go
func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)
}

func init() {
	fmt.Println("we can have more than one func init()")
}

//the _, is a pointer to a file
func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {

		//fmt.Println("error happened", err) //error happened open no-file.txt: The system cannot find the file specified.
		log.Println("error happened", err)
		//log.Fatalln(err)
	}
}

//println formats using the default formats for its operands and
//write to standard output

/*
package log implements a simple logging package
it writes to standard error and prints the data/time of each logged message
*/
