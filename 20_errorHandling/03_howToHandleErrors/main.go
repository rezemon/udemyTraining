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
		//log.Println("error happened", err)
		//log.Fatalln(err)
		panic(err)
	}
}

/*
panic func call panic after writing the log message
panic is good if we want to see stack stuff

*/

/*
the fatal function call os.Exit(1) after writing the log message
Fatalln is equivalent to println() followed by a call to os.Exit(1)
*/

//println formats using the default formats for its operands and
//write to standard output

/*
package log implements a simple logging package
it writes to standard error and prints the data/time of each logged message
*/
