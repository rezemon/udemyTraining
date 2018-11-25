/*
if given a link

https://godoc.org/github.com/dustinkirkland/golang-petname
package petname
import "github.com/dustinkirkland/golang-petname"

copy the ..
github.com/dustinkirkland/golang-petname

go to git bash window, type in below to exec
go get -u github.com/dustinkirkland/golang-petname

if the gopath is set to
set GOPATH=C:\goworkspace

then you can see the following in the c:\goworkspace
you should be able to see this created in the following path
/c/goworkspace/src/github.com/dustinkirkland

next is how to import and use it
*/

package main

import (
	"fmt"

	"github.com/GoesToEleven/udemyTraining/01_rezemon/mycalculator"
	// below is a super solid NAMING space
	// i understand it IMMEDIATELY, and super easy to use
	// how to use the imported function, look at the source code documentation
	"github.com/dustinkirkland/golang-petname"
	//looking at it , the package is called "petname"
)

func main() {
	fmt.Print("how to consume a github shared source")
	intA := 3
	intB := 6
	resultA := 0
	resultA = mycalculator.Myadd(intA, intB)
	fmt.Println("result of ", intA, " + ", intB, " = ", resultA) //18, because i double it

	var strA string
	strA = petname.Name()
	fmt.Println("random name of pet = ", strA)
	strB := petname.Generate(4, "#")
	fmt.Println(strB)
	strC := petname.Generate(7, "#")
	fmt.Println(strC)
	strD := petname.Generate(4, "#")
	fmt.Println(strD)

	for i := 0; i <= 10; i++ {
		fmt.Println(petname.Generate(1, "#"))
	}
}
