package main

import(
  "fmt"
  "encoding/json"
  "strings"
)

type Person struct{
  Surname string
  Name string
  Age int
  notExported int
}

func main()  {
  var p1 Person
  fmt.Println(p1)
  fmt.Println("================")
  // rdr == reader
  //below simulate something coming in from internet
  //for this tutorial, we simulate it this way
  //REZE
  //this also another method to simulate something coming in from internet
  rdr:=strings.NewReader(`{"Surname":"lee","Name":"chuan boon","Age":33}`)
  json.NewDecoder(rdr).Decode(&p1)
  //Decode reads the next JSON-encoded value from its
  //input and stores it in the value pointed to by v
  //reze: because it pointed to, meaning the address
  // has to decode to the Address

  fmt.Println(p1)
  fmt.Println(p1.Surname)
  fmt.Println(p1.Name)
  fmt.Println(p1.Age)
}
