//provide more information , instead of just string of error text
/*
######################################################
i guess, this context will be what i used in the debug
######################################################

in the fmt package
there is a
func Errorf() error, and there is empty interface{}
:), accept all types
*/

package main

import (
	"fmt"
	"log"
	//"github.com/GoesToEleven/udemyTraining/_myTools/Describe"
)

//put Err infront
//it is idiomatic, so intellisense will help

func main() {

	_, err := sqrt(-10.233)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(aNum float64) (float64, error) {
	if aNum < 0 {
		//now has more information, this could be use for debugging
		return 0, fmt.Errorf("norgate math again: square root of negative number : %v", aNum)
	}
	//implmentation
	return 42, nil
}
