package main

import (
	"fmt"
	"sort"

	"github.com/GoesToEleven/udemyTraining/_myTools/Describe"
)

type aIntSlice []int

/*
######################
REASON FOR ALIAS !!!
######################
remember the tutorial say NO MEANING to make a alias of type int
e.g. type myIntType int
UNLESS UNLESS UNLESS
we need to attach a method to it
and in order to work with interface, we need method to attach to the type
hence, we have to make aliase of the type we want to work with
*/
func (a aIntSlice) Len() int {
	return len(a)
}

func (a aIntSlice) Less(i int, j int) bool {
	return a[i] < a[j]
}

func (a aIntSlice) Swap(i int, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

/* DONT WORK
func (nnn []int) Len() int {
	return len(nnn)
}
*/

func main() {
	//var nnn []int
	//sort.Sort(nnn)
	//sort.Sort([]int{4, 3, 2, 1})
	fmt.Println("learning interface")
	var nnn aIntSlice // aaa is of type []int
	nnn = []int{22, 13, 16, 3, 88, 6, 2, 10}
	sort.Sort(nnn)
	myTools.Describe(nnn)
	nnn2 := sort.Reverse(nnn)
	myTools.Describe(nnn2)
	sort.Sort(nnn2)
	myTools.Describe(nnn2)
}
