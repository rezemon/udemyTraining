//some queries that i have and try it out here
//http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
//http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

package main

import(
  "fmt"
)

//declare a interface that is always able to accept any interface
/*
most of the time
type Animal interface{
  Name string // a field or a properties
  Speak() string //method
}
by Duck Typing, if it QUACKS, it is a Duck
type SomeInterface interface{}
this is a EMPTY interface, and it is always CORRECT
every type in GOLANG, implement this interface
because interface definition is MUST implement ALL methods defined in the interface
but EMPTY interface has no methods
thus
aString string can be said to implement Empty interface
aInteger int also implement Empty interface
AS IT MEETS THE CONDITION REQUIRED BY INTERFACE
 */
//the interface{} is

func DoSomething (aInterface  interface{}){
  fmt.Println("accepting a param of type interface ")
  fmt.Printf("%T", aInterface)
  fmt.Println()
}

func main(){
 //declare a interface and pass to DoSomething

 //var AA interface{} //var AA of type empty interface
 var BB interface{} //it must be a empty interface
 fmt.Printf("%T",BB) //empty interface is of type ="nil"
 fmt.Println()
// describe(BB)
 DoSomething(BB)

 var CC interface{
   aMethod() string //aMethod that accept nothing but return a string
 }

 DoSomething(CC)

var DD int
 DoSomething(DD)

 var EE string
 DoSomething(EE)


 /*
 type Animal interface {
 	Speak() string
 }

 type Dog struct {
 }

VERY VERY IMPORTANT
we create a type Dog, based on a struct
and when type Dog implment Speak() , type Dog is a ANIMAL interface , as it implement a ANIMAL
 func (d Dog) Speak() string {
 	return "Woof!"
 }

  dog is a type
  int is a also a type
  string is also a type
  because ..
  func(aInt int){
  // dont have to implement ANY method to be a EMPTY interface
}
aInt of type int, implment nothing... will be a EMPTY interface ,
it MEETS THE REQUIREMENT because EMPTY interface dont need any method to be implemented
that is why all the above wont have error

AND AND AND by parallel type
the empty interface convert itself to become WHATEVER TYPE that was passed in as param

  */

}
