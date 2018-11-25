package main

import "fmt"

func main() {

	/*
		TUTORIAL describe this as a slice[]
		containing []string
		i dont really agree
		to me, it is just multi-dimension of string
	*/
	var myStudentRecord = make([][]string, 0)

	var aStudent1 = make([]string, 4)
	aStudent1[0] = "lee"
	aStudent1[1] = "chuan boon"
	aStudent1[2] = "88.88"
	aStudent1[3] = "99.99"

	myStudentRecord = append(myStudentRecord, aStudent1)

	var aStudent2 = make([]string, 4)
	aStudent2[0] = "tan"
	aStudent2[1] = "kim choo"
	aStudent2[2] = "88.00"
	aStudent2[3] = "99.00"

	myStudentRecord = append(myStudentRecord, aStudent2)

	fmt.Println("my student record = ")
	fmt.Println(myStudentRecord)
}
