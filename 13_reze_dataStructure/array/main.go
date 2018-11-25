package main

import "fmt"

func main() {

	var aArray [10]string

	fmt.Println("length of aArray = ", len(aArray))

	fmt.Print("array content = ")
	fmt.Println(aArray)
	fmt.Print("type of aArray = ")
	fmt.Printf("%T", aArray)
	fmt.Println()

	aArray[2] = "lee chuan boon"
	fmt.Print("array content = ")
	fmt.Println(aArray)
	aArray[3] = "tan kim choo"
	fmt.Println(aArray)
	aArray[9] = "this is good"
	fmt.Println(aArray)
	aArray[6] = "P6" //position 7, as array start with index ZERO
	fmt.Println(aArray)

	//golang is very suitable for me
	//seems to be what i was doing 20 years ago, or 30 years ago
	//else it would not be so easy for me
	//now try 2 dimension array without seeing the tutorial

	//var multiDimArray [2][2]int //2 x2 means 3 x 3 as array index start with zero
	//wrong should use 3 x 3
	var multiDimArray [3][3]int
	fmt.Println("length of multiDimArray = ", len(multiDimArray))

	fmt.Print("array content = ")
	fmt.Println(multiDimArray)
	fmt.Print("type of multiDimArray = ")
	fmt.Printf("%T", multiDimArray)
	fmt.Println()
	//want to put it with 1,2,3 ... 9
	//remember INDEX is still ZERO
	k := 1

	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			multiDimArray[i][j] = k
			k++
		}
	}
	fmt.Println(multiDimArray)
}
