package main

import(
  "fmt"
)

type Person struct{
  surname string
  name string
  age int
}

func (p Person)Greeting(){
  fmt.Println("i am a normal person")
}
type SecretAgent struct{
  Person
  licenseToKill bool
  age int //will this conflict
}

func (sa SecretAgent)Greeting(){
  fmt.Println("i am a secret agent")
}
func main()  {
  person1 := SecretAgent{
    Person : Person{
      surname:"lee",
      name:"chuan boon",
      age:33,
    },
    licenseToKill:true,
    age:19,

  }

person2 :=SecretAgent{
  Person : Person{
    name:"kim choo",
    age:32,
    surname:"tan",
  },
  licenseToKill:false,
  age:18,
}
  fmt.Println(person1)
  fmt.Println(person2)
  fmt.Println(person1.age,person1.licenseToKill)


  fmt.Println(person2.age) //outer age
  fmt.Println(person2.Person.age) //inner age

  person2.Greeting() //outer person
  person2.Person.Greeting() // inner person //(embedded person)

}
