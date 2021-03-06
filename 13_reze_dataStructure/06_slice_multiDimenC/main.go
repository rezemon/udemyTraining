package main

import (
	"fmt"

	"github.com/dustinkirkland/golang-petname"
)

func main() {

	var phoneBook = make([][]string, 0) //we dont know the size, we dont state the capacity

	// assume to add 3 phone entries into it
	// phoneBook has a list of 3 entries
	// each entry is also a slice
	// slice of {name, phoneNumber}

	//we need a func to return a random name,
	//search for one, or create one
	// temporary, we use the index ??

	for i := 0; i <= 2; i++ {
		var phoneEntry = make([]string, 2)
		phoneEntry[0] = string(i + 65) //this become capital A ??
		phoneEntry[1] = petname.Generate(1, "")

		phoneBook = append(phoneBook, phoneEntry)
	}

	fmt.Println(phoneBook)
	// interesting, above works!!!
	//add more using another method

	for i := 0; i <= 2; i++ {
		//we give it another variable name, even though reusing phoneEntry is ok
		//as the scope is within the CURLY brackets only
		var phoneEntry2 = make([]string, 0) //,0 means slice size can be expanded
		// we store 3 random string to phoneEntry2
		for j := 0; j <= 2; j++ {
			phoneEntry2 = append(phoneEntry2, petname.Generate(1, "#"))
		}
		phoneBook = append(phoneBook, phoneEntry2)
	}

	fmt.Println(phoneBook)
	// above interesting, show the slice contain slice of different info

	for i := 0; i <= 2; i++ {
		//we give it another variable name, even though reusing phoneEntry is ok
		//as the scope is within the CURLY brackets only
		//phoneEntry2 is localized,
		var phoneEntry2 = make([]string, 0) //,0 means slice size can be expanded
		// we store 3 random string to phoneEntry2
		for j := 0; j <= 2; j++ {
			phoneEntry2 = append(phoneEntry2, petname.Generate(3, "#"))
		}
		phoneBook = append(phoneBook, phoneEntry2)
	}

	fmt.Println(phoneBook)

	student := make([]string, 0)
	students := make([][]string, 0)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println((student == nil))

	var studentB []string
	fmt.Println(studentB == nil)

	studentC := []string{}
	fmt.Println(studentC == nil) //suppose to be nil according to tutorial
	//make establish the POINTER to the underlying array
	//thus for make it is NOT nil
	// though it is not nil
	// UNLIKE make, it did not point to the underlying array
	//if you perform an index reference , you get an error
	//but if you APPEND to it, it is fine
	//golang will know that you want to put something
	//and intelligently will create the underlying array
	//and append the first item to it

	//studentD := []string //ERROR NOT ALLOWED
	// fmt.Println((studentD == nil))

}
