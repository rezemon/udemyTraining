package main

import (
	"fmt"
)

func main() {

	var aMap = make(map[string]string)

	aMap["lee"] = "chuan boon"
	aMap["tan"] = "kim choo"
	aMap["wong"] = "keng meng"

	fmt.Println("aMap = ", aMap)

	aValue, present := aMap["tan"]
	fmt.Println("aValue = ", aValue)
	fmt.Println("present = ", present)

	aValue2, present2 := aMap["tantan"]
	fmt.Println("aValue2 = ", aValue2)
	fmt.Println("present2 = ", present2)

	//map dont have append function, so create map with MAKE !!!
	//otherwise trapped,
	fmt.Println("length of aMap =", len(aMap))

	aMap["test"] = "to be deleted"
	fmt.Println(aMap)

	fmt.Println("deleting the element with the key 'test")

	delete(aMap, "test2") //if we specify a non-existing key, NOTHING happen, no error triggered
	fmt.Println(aMap)

	delete(aMap, "test")
	fmt.Println(aMap)

}
