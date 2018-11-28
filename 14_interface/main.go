package main

import(
  "fmt"
)

type Shape interface{
  Area(aFloat float64) (float64)
}

type Square struct{
  side float64
  }

//in order for Square to behave like a Shape
//it has to implement the Area method
func (sq Square) Area(aFloat float64)float64{
  return (sq.side * sq.side)
}

func (sq Square) changeSideValue(aFloat float64){
  sq.side = aFloat
}
//this inteface Shape CAN ACCEPT CAN ACCPET CAN ACCEPT
//any interface TYPE that implement the method Area
func getArea(aInterface Shape) {
  fmt.Println("the shape is = ",aInterface )
  aInterface.Area(6)
}
func main()  {
  var aSquare Square
  aSquare.side = 8.8
  aArea := aSquare.Area(aSquare.side)
  fmt.Println("aArea = ", aArea)
  aSquare.side = 6.0
  aArea = aSquare.Area(aSquare.side)
  fmt.Println("aArea = ", aArea)

  aSquare.changeSideValue(3)
  //the square here is passed in as value
  //hence no changes took place, it still remain as 6.0 
  aArea = aSquare.Area(aSquare.side)
  fmt.Println("aArea = ", aArea)
 //above value not affected, as passed in value

}
