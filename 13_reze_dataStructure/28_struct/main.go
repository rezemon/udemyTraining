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
  Name string `json:"-"` // telling not to include
  Age int `json:"wisdom score"` //telling Age to change to wisdom score
}

func main()  {
  p1:=Person{"lee","chuan boon",33}
  bs,_ := json.Marshal(p1)
  fmt.Println(string(bs))
  p2:=PersonB{"wong","ken meng",34}
  bs2,_:= json.Marshal(p2)
  fmt.Println(string(bs2))
}
