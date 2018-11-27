package main

import(
  "encoding/json"
  "fmt"
)

type Person struct{
  Surname string
  Name string
  Age int
  notexported int //because it is in small caps
}
func main()  {
  p1:= Person{"lee","chuan boon",33,007}
  //bs = byte stream
  bs,x := json.Marshal(p1)
  fmt.Println("x = ", x)
  fmt.Print("x type = ")
  fmt.Printf("%T",x)
  fmt.Println()

  fmt.Println("bs = " , bs)
  fmt.Print("bs type = ")
  fmt.Printf("%T",bs)
  fmt.Println()

  //and this is JSON
  //from a STRUCT change to JSON
  //marshal, is from a STRUCT change to JSON
  //and this work on a string 
  fmt.Println("string representation of bs =",string(bs))

}
