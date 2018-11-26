package main

import(
  "fmt"
)

//Animal what is a animal
type Animal interface {
  Speak() string
  AssignName(aName string) string
}

type Dog struct{
  Name string
}

func (d *Dog)Speak() string{
  aString := "Woof Woof"
  return aString
}

func(d *Dog)AssignName(aDogName string) string{

  d.Name = aDogName
  aString := "dog name is now = " + aDogName
  return aString
}

type Cat struct{
  Name string
  }

  func (c *Cat)Speak() string{
    aString :="Meow Meow"
    return aString
  }

/*
https://tour.golang.org/methods/8
https://tour.golang.org/methods/8
*/

/*
Choosing a value or pointer receiver
There are two reasons to use a pointer receiver.
The first is so that the method can modify the value that its receiver points to.
The second is to avoid copying the value on each method call.
This can be more efficient if the receiver is a large struct,
for example.
In this example, both Scale and Abs are with receiver type *Vertex, even though the Abs method needn't modify its receiver.
In general, all methods on a given type should have
either value or pointer receivers,
but not a mixture of both. (We'll see why over the next few pages.)

PUT a *Cat , make
 */
//  func(c Cat)AssignName(aCatName string) string{
func(c *Cat)AssignName(aCatName string) string{
    c.Name = aCatName
    aString := "cat name is now = " + aCatName
    return aString
  }


type Me struct {
  Name string
}

func (m *Me)Speak()string{
  aString := ("Yeap Yeap Yeap")
  return aString
}

func(m *Me)AssignName(aPersonName string) string{
  m.Name = aPersonName
  aString := "person name is now = " + aPersonName
  return aString
}

type Bird struct{
  Name string
}

func (m *Bird)Speak()string{
  aString := ("Chirp Chirp Chirp")
  return aString
}


func main(){
  fmt.Println("learning interface ")

  aCat := &Cat{"Garfield"}
  fmt.Println(aCat.Speak())
  fmt.Println("cat name = ", aCat.Name)
  fmt.Println(aCat.AssignName("Cat Women"))
fmt.Println("cat name =", aCat.Name)
aCat.Name="AAA"
fmt.Println("cat name =", aCat.Name)

//this way, regardless of what
//
//create a slice containing Animal interface
aSlice := make([]Animal,0)
//add or rather append the interface to it
//what is a Animal interface
/*
as long as
instead of designing our abstractions in terms of what kind of data
our types can hold, we design our abstractions in terms of
what actions our types can execute.

https://gobyexample.com/interfaces
The circle and rect struct types
both implement the geometry interface
so we can use instances of these structs as arguments to measure.

 */
/*
# github.com/GoesToEleven/udemyTraining/13_reze_dataStructure/20_interface
.\main.go:109:16: cannot use aDog (type Dog) as type Animal in append:
	Dog does not implement Animal (AssignName method has pointer receiver)

aDog := Dog{"Cupcake"}
SOLUTION SOLUTION SOLUTION, put a & to make aDog a pointer
then we have implement

 */
aDog := &Dog{"Cupcake"} //wow, now a dog is a pointer ???
aMe := &Me{"lee chuan boon"}
aSlice = append(aSlice, aCat)
aSlice = append(aSlice, aDog)
aSlice = append(aSlice, aMe)

for index,value := range aSlice{
  //the value here is a INTERFACE, it is an ANIMAL interface
  //because it implement all methods of ANIMAL interface
  //value being a ANIMAL interface, automatically intellisense the methods
  //
 fmt.Println("index =",index)
  fmt.Println(value.Speak())
}

//if we create another Animal, but did not implement one of the method
//the compiler should complain, cannot accept
aBird := &Bird{"Tweety Bird"}
fmt.Println(aBird.Speak()) // this still work!!!
//BUT BUT BUT aBird is STRICTLY NOT A Animal
//AS IT DID NOT IMPLEMENT THE ANIMAL INTERFACE FULLY

/*
BELOW ERROR
github.com/GoesToEleven/udemyTraining/13_reze_dataStructure/20_interface
.\main.go:146:16: cannot use aBird (type *Bird) as type Animal in append:
	*Bird does not implement Animal (missing AssignName method)
 */

// aSlice = append(aSlice,aBird)

//because i implement Dog, Cat, Me as pointer receiver
//because it is a pointer receiver
//changes i make will
//The first is so that the method can modify the value that its receiver points to.
aMe.AssignName("Rezemon Lee")
fmt.Println("alias for lee chuan boon", aMe.Name)

//instead of using a Slice of ANIMAL INTERFACE
//we could define a func
//to demo the abstractions
//it will just call method of Speak()
//but today Dog speak "Woof Woof"
//future changes in implementation
//Dog can speak "Bark Bark"

AnimalSound(aDog) //aDog is a pointer , the implementation is a POINTER RECEIVER
AnimalSound(aMe)
}

//a function that has param of "Animal interface"
func AnimalSound(aAnimal Animal){
  aString := aAnimal.Speak()
  fmt.Println(aString)
}
