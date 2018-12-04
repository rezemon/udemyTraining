package main

import (
	"fmt"

	"github.com/GoesToEleven/udemyTraining/_myTools/Describe"
)

func main() {

	output := func(aInt int) {
		aInt++
		fmt.Println("the integer value = ", aInt)
	}

	output2 := func(aInt int) string {
		aInt++
		return ("i am good")
	}

	//it is a pointer to a func(int)
	myTools.Describe(output) //(0x490150, func(int))

	//what is type of output2
	myTools.Describe(output2) //(0x490220, func(int) string)
}
