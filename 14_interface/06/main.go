package main

import(
  "fmt"
  "github.com/GoesToEleven/udemyTraining/_myTools/Describe"
  "math"
)

type Shape interface{
  area() float64
}

type Square struct {
  side float64
}

type Circle struct {
  radius float64
}

func (sq Square) area() float64{
  return (sq.side * sq.side)
}

func (c Circle) area() float64{
  return (math.Pi * c.radius *c.radius)
}

func main()  {
  fmt.Println("learning interface")
  aSquare := Square{3}
  result := shapeArea(aSquare)
  myTools.Describe(result)

  aCircle := Circle{2}
  result = shapeArea(aCircle)
  myTools.Describe(result)

}

func shapeArea(aShapeInterface Shape) float64{
  return aShapeInterface.area()
}
