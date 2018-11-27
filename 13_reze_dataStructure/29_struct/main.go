package main

import(
  "fmt"
  "encoding/json"
)

type Person struct{
  Surname string
  Name string
  Age int
}

type PersonB struct{
  Surname string
  Name string
  Age int `json:"wisdom score"` //adding this tag, now it is imported
}

func main()  {
  //we receive a JSON, and convert it to STRUCT
  var p1 Person
  fmt.Println(p1)
  fmt.Println(p1.Surname)
  fmt.Println(p1.Name)
  fmt.Println(p1.Age)

//below is the JSON, could be from a API sending us JSON
//and we need to convert it back to a STRUCT
//need ``because there is double quote, so single code need to be used
//here i take a STRING
//and changed it to slice of byte
//changed it to slice of byte
  bs:= []byte(`{"Surname":"lee","Name":"chuan boon","Age":33}`)

  json.Unmarshal(bs, &p1)
  //fmt.Println(p1)
  fmt.Println("###")
  fmt.Println(p1.Surname)
  fmt.Println(p1.Name)
  fmt.Println(p1.Age)
  fmt.Printf("%T", p1)
  fmt.Println()

  var p2 PersonB
// take note, we are expecting field Age ,
// but what is send to us is "wisdom score"
// if we dont use tag to handle it, what will happen
  bs2:= []byte(`{"Surname":"wong","Name":"ken meng","wisdom score":34}`)
// the Age is not imported, so need a tag

  json.Unmarshal(bs2,&p2)

  fmt.Println(p2)
}
