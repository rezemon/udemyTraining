package main

import(
  "fmt"
)

type Person struct{
  surname string
  name string
  age int
}

func (p Person) fullname()string{
  return p.surname + p.name
}
func main(){
  person1 := Person{"lee","chuan boon",33}
  person2 := Person{"tan","kim choo",32}
  fmt.Println(person1.fullname())
  fmt.Println(person2.fullname())
}
