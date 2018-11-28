package myTools

import(
  "fmt"
)


//a empty interface{}, means can take any value
type I interface{
  // I is empty interface
  // no methods
  // thus all TYPE are empty interface
}

/* remove comments to test this func Describe
func main()  {
  //https://godoc.org/golang.org/x/tools
  //https://godoc.org/golang.org/x/tools
  fmt.Println("how to use describe")
  aInt := 33
  describe(aInt)
}
*/
func Describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)

}
