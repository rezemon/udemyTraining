package main

import(
  "fmt"
  "sort"
  "github.com/GoesToEleven/udemyTraining/_myTools/Describe"

)

type aIntSlice []int

func main()  {
  fmt.Println("learning interface")
 var n1 aIntSlice
  n1 = []int {4,3,2,1,8}

fmt.Println(n1)

// this sort.Interface is another thing..
//var sortInterface sort.Interface
//fmt.Println(sortInterface.Len())

//https://godoc.org/sort#IntSlice
//looking at type IntSlice
var myIntSlice sort.IntSlice
myIntSlice = []int {12,3,5,62,1,7,88,2}
myIntSlice.Sort()
fmt.Println(myIntSlice)
//how to do reverse
//var myReverseIntSlice sort.Interface

// sort.Reverse(myIntSlice) return a Interface
aSomething := sort.Reverse(myIntSlice) //a something is a interface
// interface is ??
//reason that we can pass myIntSlice into Reverse(data Interface)
/*
because to be a Duck, have quack like a Duck
To be Duck, behave all things required by a Duck
to be a Inteface(cap I), need to implement func Len(), func Less(i,j int), func Swap(i,j)
myIntSlice which is of type sort.IntSlice
https://godoc.org/sort#IntSlice
type IntSlice []int
implement
func (p IntSlice) Len() int
func (p IntSlice) Less(i,j int) bool
func (p IntSlice) Swap(i,j)
thus
IntSlice is BEHAVING EXACTLY LIKE Interface(cap I)
https://godoc.org/sort#IntSlice
hence
func Reverse
func Reverse(data Interface) Interface
we therefore can pass IntSlice as a param
we therefore can pass myIntSlice as a param
because
var myIntSlice sort.IntSlice

 */

//myTools.Describe(&aSomething)
fmt.Println("===============")
myTools.Describe(aSomething)
fmt.Println("===============")
fmt.Println(aSomething.Len())
sort.Sort(aSomething) //aSomething will be reverse when we sort it
myTools.Describe(aSomething)
sort.Sort(aSomething)
myTools.Describe(aSomething)

var myStringSlice sort.StringSlice
sss :=[]string{"Zeno","John","Al","Jenny"}
myStringSlice = sss
myStringSlice.Sort()
myTools.Describe(sss)

type people []string
studyGroup := people{"Zebra","Jaguar","Ant","Joy"}

myTools.Describe(studyGroup)
convertedStudyGroup := []string(studyGroup)
myStringSlice = convertedStudyGroup
myStringSlice.Sort()
myTools.Describe(myStringSlice)
myTools.Describe(convertedStudyGroup)
}
