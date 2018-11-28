package main

import(
  "fmt"
  "github.com/GoesToEleven/udemyTraining/_myTools/Describe"
  "math"
)

// able to RETURN any type using the empty interface{}
type Shape interface{
  //area() float64
  area() interface{}
}

type Square struct {
  side int
}

type Circle struct {
  radius float64
}

type Unknown struct{
  xFactor string
}

func (sq Square) area() interface{}{
  return (sq.side * sq.side)
}

func (c Circle) area() interface{}{
  return (math.Pi * c.radius *c.radius)
}

func (u Unknown) area() interface{}{
  return "X files"
}

func main()  {
  fmt.Println("learning interface")
  aSquare := Square{3}
  result := shapeArea(aSquare)
  myTools.Describe(result)

  aCircle := Circle{2.34}
  result = shapeArea(aCircle)
  myTools.Describe(result)

  aUnknown := Unknown{"anything"}
  result = shapeArea(aUnknown)
  myTools.Describe(result)
}

func shapeArea(aShapeInterface Shape) interface{}{
  return aShapeInterface.area()
}
