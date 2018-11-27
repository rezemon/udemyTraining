package main

import(
  "fmt"
)

type Animal interface{
  Speak() string
}

type Dog struct{

}

//aDog implement the Animal interface
//Speak method is with receiver type *Dog, which is a POINTER RECEIVER
func (d *Dog) Speak() string {
  return "Woof Woof"
}

type Bird struct{

}

//aDog implement the Animal interface
//Speak method is a value type(Bird), it is NOT a pointer receiver
func (d Bird) Speak() string {
  return "Chirp Chirp"
}


type Cat struct{
  // dont implement the Speak(), so it is not a ANIMAL
}
func main(){
  //testing the interface type
  fmt.Println("what is the type")

  //var aDog Dog
  aDog := &Dog{} //must be this , Dog receiver is a pointer
  fmt.Printf("%T",aDog) //main.Dog
  fmt.Println()
  fmt.Println(aDog.Speak())

  var aBird Bird //different from a Dog, Bird receiver is
  fmt.Printf("%T",aBird) //main.Bird
  fmt.Println()
//how to assign Animal interface to a Dog
//one method by passing as param in a method
animalSpeak(aDog)
animalSpeak(aBird)
var aCat Cat           //comment this if want this program to work
animalSpeak(aCat)      // comment this if want this program to work
/*
# github.com/GoesToEleven/udemyTraining/13_reze_dataStructure/22_interface
.\main.go:53:12: cannot use aCat (type Cat) as type Animal in argument to animalSpeak:
	Cat does not implement Animal (missing Speak method)
 */
}

// this is a function expecting a animal interface
//if we pass in a Dog , it is ok, if Dog already implement what Animal need
//if we pass in a Cat , we should see an Error
func animalSpeak(aAnimal Animal){
  fmt.Println(aAnimal.Speak())
}
