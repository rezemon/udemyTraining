package main

import(
  "fmt"
  "sort"
)

/*
https://godoc.org/sort#IntSlice
we look at
type IntSlice []int
and IntSlict implemented func Len, func Less, func Search
so IntSlice is a Interface  (with capital I)
therefore, we can just call it as long as it is []int

type Interface ¶ Uses
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
*/

type n1 struct{
  n1slice []int
}
func (n n1) Len() int{
  return len(n.n1slice)
}

func (n n1) Less( i int , j int)bool{
  if (i < j){
    return true
} else {
  return false
}
}

func (n n1) Swap(i int, j int){
  //find the index first
  //how to find index
  //scan ??
  var iIndex int
  var jIndex int
 for index,value := range n.n1slice{
   //fmt.Println(index, " - ", value) //dont print first
   if value == i {
     iIndex = index
   }
   if value ==j {
     jIndex = index
   }
 }

//swapping
   temp:= n.n1slice[iIndex]
   n.n1slice[iIndex ] = n.n1slice[jIndex]
   n.n1slice[jIndex] = temp

}
func main()  {
  fmt.Println("learning interface")
  //[]int{22,67,2,1,17,7,10,8}

  var n2 n1
  n2.n1slice = []int{22,67,2,1,17,7,10,8}

for index,value := range n2.n1slice{
  fmt.Println(index, " - ",value)
}
  //sort.Sort(&n2)
  sort.Sort(n2)
  fmt.Println(n2.n1slice)
}
