package main

import(
  "fmt"
  "github.com/GoesToEleven/udemyTraining/_myTools/Describe"
)

type Shape interface{
  area(aInt int) int
}

type Square struct{
  side int
}

func (s Square) area (sInt int) int{
  return s.side * s.side
}
type Circle struct{
  side int
}

func (c Circle) area (cInt int) int{
  return c.side * 10
}

func main()  {
  fmt.Println("learning interface")
  aSquare := Square{3}
  var aTempInt  int
  aTempInt = shapeArea(aSquare)
  myTools.Describe(aTempInt)
}

//this wont work, we need to use pointer to Struct
func shapeArea(aShapeInterface Shape) int{
  return aShapeInterface.area(0)
}
