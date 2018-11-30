//method can be attached to any name type
package main

import (
	"fmt"
	"sort"
)

type ttt []string

var sss = []string{"zebra", "john", "alvin", "jenny"}

func (t ttt) Hello() {
	fmt.Println("i am user defined type ttt")
}

/*
sss is not a type

func (s sss) Hello() {
	fmt.Println("does this work")
}
*/
func main() {
	var aType ttt
	aType.Hello()

	var aaa sort.StringSlice
	aaa = []string{"zebra", "john", "alvin", "jenny"}
	sort.Sort(aaa)
	fmt.Println(aaa)
	//aaa is already implementing Interface interface  of sort package
	//sort.Sort
	//sort. in the Package
	//.Sort is the func
	bbb := []string{"zoo", "japan", "australia", "josephine"}
	//convert it
	ccc := sort.StringSlice(bbb)
	sort.Sort(ccc)
	fmt.Println(ccc)
	// https://golang.org/pkg/sort/#sort

	var nnn []int
	nnn = []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Sort((sort.IntSlice(nnn)))
	fmt.Println(nnn)

	sort.Sort(sort.Reverse(sort.IntSlice(nnn)))
	fmt.Println(nnn)

}
