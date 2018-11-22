package main

//"C:/goworkspace/src/github.com/GoesToEleven/udemyTraining/01_rezemon/myCalculator"
import (
	"fmt"

	"github.com/GoesToEleven/udemyTraining/01_rezemon/mycalculator"
)

func main() {
	var aNumber int
	aNumber = 88
	fmt.Println("number = ", aNumber)

	var numberA int
	var numberB int
	var resultC int
	numberA = 11
	numberB = 22
	resultC = mycalculator.Myadd(numberA, numberB)
	//resultC = myAdd(numberA, numberB)
	fmt.Println("result =", resultC)

	fmt.Println(mycalculator.MyName)
	//fmt.Println(mycalculator.myName) //Unexported , small caps means not visible OR not exported

}
