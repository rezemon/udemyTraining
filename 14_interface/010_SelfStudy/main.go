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

func (s *Square) area (sInt int) int{
  s.side = sInt
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
  //var aSquare *Square
  //aTempSquare := Square{3}
  aSquare := Square{2}
  fmt.Println(aSquare.area(aSquare.side))

  var aTempInt  int
  aTempInt = shapeArea(&aSquare, 3)
  myTools.Describe(aTempInt)

  aTempInt= aSquare.area(4)
  myTools.Describe(aTempInt)

  aTempInt= shapeArea(&aSquare,5)
myTools.Describe(aTempInt)

}

//this wont work, we need to use pointer to Struct
func shapeArea(aShapeInterface Shape, aShapeValue int) int{
  return aShapeInterface.area(aShapeValue)
}
