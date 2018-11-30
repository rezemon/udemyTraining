//tutorial
//https://godoc.org/sort#Interface
//referring to the package sort
//figure out how to sort

package main

import(
  "fmt"
  "github.com/GoesToEleven/udemyTraining/_myTools/Describe"
  //"sort"
)

func main()  {
fmt.Println("learning interface")
  sss:= []string{"Zeno","John","Alvin","Jenny"}

  nnn:= []int{7,4,8,1,9,19,12,32,3}
  type people []string

  studyGroup := people{"Zeno","John","Alvin","Jenny"}

type aa []int //aa is a type

//aa =1,2,3,4
var bb aa //bb is of type aa
var cc aa
var dd aa
var ee aa

  myTools.Describe(sss)
  myTools.Describe(nnn)
  //myTools.Describe(people)
  myTools.Describe(studyGroup)
  myTools.Describe(bb)
  aSliceOfTypeAa := []aa{cc,dd,bb}
  myTools.Describe(aSliceOfTypeAa)
  bb = []int{1,2,3,4,88,99}
  cc= []int{8,7,6}
  myTools.Describe(aSliceOfTypeAa)
  myTools.Describe(bb)
  myTools.Describe(cc)
  myTools.Describe(dd)
  fmt.Println(aSliceOfTypeAa)

  ee =[]int{77,88,99}
  aSliceOfTypeAa = append(aSliceOfTypeAa,ee)
 myTools.Describe(aSliceOfTypeAa)
 dd =[]int{66,44,22}
 myTools.Describe(aSliceOfTypeAa)

 aSliceOfTypeAa2 := []*aa{&bb,&cc,&dd,&ee}
 myTools.Describe(aSliceOfTypeAa2)
 myTools.Describe(&bb)
fmt.Println(*(&bb)) // & is the address, * is the pointer


//try sorting nnn
/*
type IntSlice
func (p IntSlice) Len() int
func (p IntSlice) Less(i, j int) bool
func (p IntSlice) Search(x int) int
func (p IntSlice) Sort()
func (p IntSlice) Swap(i, j int)
 */

/*
i have to implement the interface before i can use it
nnn is a STRUCT ??
sss is a STRUCT ???
*/
type qqq struct{

}

rrr := qqq{}
myTools.Describe(nnn)
//myTools.Describe(qqq)
//fmt.Printf("%T",qqq)
myTools.Describe(rrr) //struct of type main.qqq

//implement the interface
/*
in order to call
func Sort(data Interface)
to pass in , my datatype has to be of type Interface
that means i have to implement this interface

type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}

so []int{22,33,2,14,99,6}
*/

fmt.Println(len(nnn)) //nnn implement len() int , wrong
// Interface asking for capital

pppStruct :=[]int{7,4,8,1,9,19,12,32,3}

myTools.Describe(pppStruct)


} //End main
type ppp struct{
  pppInt []int
}

//{7,4,8,1,9,19,12,32,3}
// implemented the func Len
func(p ppp) Len() int{
   return len(p.pppInt)
}
