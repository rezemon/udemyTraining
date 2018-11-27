package main
import(
  "fmt"
)

type foo int //we create a TYPE
//it is of type int , named as foo
func main(){
  var aFooInt foo
  aFooInt = 33
   fmt.Println(aFooInt)
   fmt.Printf("%T", aFooInt) //,aain.foo
fmt.Println()

   var aInt int
   aInt = 66
   fmt.Println(aInt)
   fmt.Printf("%T",aInt)
   fmt.Println()

resultA := int(aFooInt) + aInt
resultB := aFooInt + foo(aInt)

fmt.Println(resultA,resultB)
}
