package main

import(
  "fmt"
)

type Person struct{
  surname string
  name string
  age int
}
func main()  {
  person1 := Person{
    surname:"lee",
    name:"chuan boon",
    age : 33,
  }

  person2 := &Person{
    surname:"wong",
    name:"ken meng",
    age:32,
  }

  fmt.Println(person1)
  fmt.Printf("%T",person1)
  fmt.Println()
  fmt.Println(person1.surname)

  fmt.Println(person2)
  fmt.Printf("%T",person2)
  fmt.Println()
  fmt.Println(person2.surname)

}
