package  main

import (
  "fmt"
)

func main(){
  runeletter := 'A' //rune
  fmt.Println(runeletter)
  fmt.Printf("%T",runeletter)
  fmt.Println()
  // runeletterTWO := 'AB' //rune //rune is a single character
  //fmt.Println(runeletterTWO)
  // fmt.Printf("%T",runeletterTWO)

  realletter := "A" //real letter //double quotes
  fmt.Println(realletter)
  fmt.Printf("%T",realletter)
  fmt.Println()
  realletterTWO := "AB" //rune
  fmt.Println(realletterTWO)
  fmt.Printf("%T",realletterTWO)
  fmt.Println()

  runeletterTWO := rune("HABC or anything"[0]) //72
  fmt.Println(runeletterTWO)
  fmt.Printf("%T",runeletterTWO)
  fmt.Println()

  for i:=65;i<=122;i++{
    fmt.Println(i, " - ", string(i), " - ", i%12)
  }


}
