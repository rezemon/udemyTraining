/*
i want to learn about the pointer and address
have a few example
so to clear up some concepts

this demo a field pointer
next i want to learn func pointer ??
is there such a pointer

printing all these too suffering, i will have to use "describe". some package in the golang 
 */
package main

import(
  "fmt"
)

func changeValue(something int){
something++
}

//something is a int pointer
//int pointer pointing to "something"
func changeValuePointer(something *int){
  fmt.Print("type of something = ")
 fmt.Printf("%T",something) //*int
 fmt.Println()
 fmt.Println("address value of something = ",something)
fmt.Println("value at that address = ",*something) //print what is something pointing to
//*something is for HUMAN
//&something is for COMPUTER
*something++

}
func main()  {
  var aSomething int
  aSomething= 3
  fmt.Println("orginal value = ", aSomething)
  fmt.Println("change via calling changeValue")
  changeValue(aSomething)
  fmt.Println("new value is = ",aSomething)
  fmt.Println("change via calling changeValuePointer")
  changeValuePointer(&aSomething) // WE PASS ADDRESS to POINTER
  fmt.Println("new value is = ",aSomething)
}
