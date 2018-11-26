package main

import(
  "fmt"
)

type person struct{
  surname string
  name string
  age int
}

type foo int

func main(){

  aPerson := person{"lee", "chuan boon",33}
  aPersonB := person{"tan", "kim choo", 32}

  fmt.Println("aPerson = ", aPerson)
  fmt.Println("aPersonB =", aPersonB)

  var aPersonC person
  aPersonC.age = 34
  aPersonC.surname = "wong"
  aPersonC.name ="ken meng"

  fmt.Println("aPersonC = ", aPersonC)

  //when we create a type in golang
  //that type will have a UNDERLYING TYPE

 var myAge foo
 myAge = 33 //it can take this int, because foo UNDERLYING type is "int"
 /*
 though this example a bit weird , as foo is alias of int
 just like type rune is alias for type int32
 why need to alias int when you already has int
 the reason i can think of is to make the code more readable ???
 but another reason, is we want to ATTACH a METHOD to it
 the default type  int CANNOT have method
 but the alias type foo CAN CAN CAN have method
  */
 fmt.Printf("%T",myAge)   //main.foo
 //main.foo is MY OWN TYPE, i just created my own type

/*
we can create our OWN TYPE
so instead of creating a class , like other OOP
we can create our OWN TYPE

in golang
we create a type
we create a variable of the type
this is OOP in golang,
Object oriented programming
golang has a language on its own
there is no CLASS in go
 */


}
