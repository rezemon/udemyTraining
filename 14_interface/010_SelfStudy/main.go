package main

import (
	"fmt"

	"github.com/GoesToEleven/udemyTraining/_myTools/Describe"
)

// Shape is something i will describe later
type Shape interface {
	area(aInt int) int
}

type Square struct {
	side int
}

func (s *Square) area(sInt int) int {
	s.side = sInt
	return s.side * s.side
}

type Circle struct {
	side int
}

func (c Circle) area(cInt int) int {
	return c.side * 10
}

func main() {
	fmt.Println("learning interface")
	//var aSquare *Square
	//aTempSquare := Square{3}
	aSquare := Square{2}
	fmt.Println(aSquare.area(aSquare.side))

	var aTempInt int
	aTempInt = shapeArea(&aSquare, 3)
	myTools.Describe(aTempInt)

	aTempInt = aSquare.area(4)
	myTools.Describe(aTempInt)

	aTempInt = shapeArea(&aSquare, 5)
	myTools.Describe(aTempInt)

}

//this wont work, we need to use pointer to Struct
//reze, the aShapeValue has to be sent in
//https://www.ardanlabs.com/blog/2015/09/composition-with-go.html
//https://www.ardanlabs.com/blog/2015/09/composition-with-go.html
/*
Listing 6
69 // Contractor carries out the task of securing boards.
70 type Contractor struct{}
71
72 // Fasten will drive nails into a board.
73 func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board) {
74     for b.NailsDriven < b.NailsNeeded {
75         d.DriveNail(nailSupply, b)
76     }
77 }

In listing 6, we have the declaration of the Contractor type on line 70. Again, we don’t need any state so we are using an empty struct. Then on line 73, we see the Fasten method, one of the three methods declared against the Contractor type.

The method Fasten is declared to provide a contractor the behavior to drive the number of nails that are needed
into a specified board. The method requires the user to pass as the first parameter a value
that implements the NailDriver interface. This value represents the tool the contractor will use to execute this behavior.
Using an interface type for the this parameter allows the user of the API to later create
and use different tools without the need for the API to change.
The user is providing the behavior of the tooling and the Fasten method is providing the workflow for when and how the tool is used.

my method is correct, Shape is the interface (likewise NailDriver is the interface)
nailSupply *int , b *Board are information from the Contrete type , like the Mallet and Crowbar
take note how they talk about conrete type
*/
func shapeArea(aShapeInterface Shape, aShapeValue int) int {
	return aShapeInterface.area(aShapeValue)
}
