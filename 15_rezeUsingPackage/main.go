package main

import(
  "fmt"
  "github.com/GoesToEleven/udemyTraining/01_rezemon/mycalculator"
  //C:\goworkspace\src\github.com\GoesToEleven\udemyTraining\_myTools\Describe
  "github.com/GoesToEleven/udemyTraining/_myTools/Describe"

)

func main()  {
  fmt.Println("showing how to use Package, it is something like DLL")
  aInt := 33
  myTools.Describe(aInt)
  mycalculator.Myadd(aInt, 14)

aStr := "lee chuan boon"
myTools.Describe(aStr)

var aStrPtr *string
tempStr := "wong ken meng"
aStrPtr = &tempStr
myTools.Describe(aStrPtr)
myTools.Describe(*aStrPtr)
}
