package main

import (
	"fmt"
)

//https://golang.org/doc/effective_go.html#printing
//https://golang.org/doc/effective_go.html#printing

type person struct {
	surname string
	name    string
	age     int
}

//below this work, but need to pass &person
//but it dont work by default
//to work by default, pass in the address
func (t *person) String() string {
	return fmt.Sprintf("%v/%v/%v", t.age, t.surname, t.name)
}

/*
func (t person) String() string {
	return fmt.Sprintf("%v/%v/%v", t.age, t.surname, t.name)
}
*/
func main() {

	aPerson := person{"lee", "chuan boon", 33}
	fmt.Println(aPerson)
	fmt.Println(aPerson.String())
}

//also look at the example here
//https://godoc.org/sort#
