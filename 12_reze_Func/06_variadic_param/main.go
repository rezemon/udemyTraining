package main

import "fmt"

func main() {
	/*
	   at this junction , 1,2,3,4,8,22 are just constant
	   they are not assign a type
	   so when pass to the func average, they are assumed the type float64
	*/
	aNumber := average(1, 2, 3, 4, 8, 22)

	//	aNumberTwo := average[1,3,4,6,88,17] //DONT WORK
	fmt.Println(aNumber)

	aNumberThree := average2(1, 2, 3, 4)
	fmt.Println(aNumberThree)
}

func average(aParam ...float64) float64 {
	fmt.Println(aParam)      //[1 2 3 4 8 22]
	fmt.Printf("%T", aParam) //[]float64 //this is a slice of type float64
	fmt.Println()
	return 888
}

/*
slice is just a LIST
*/

func average2(aParam ...float64) float64 {
	var total float64 //default it will be ZERO
	//but i prefer EXPLICIT
	total = 0
	//total++
	/*
			range A,
			mean it can LOOP over a list of something
			in this case, it LOOP over a list of float64
		  and it return , an INDEX and a VALUE
			because this is a LIST
	*/
	for i, v := range aParam {
		fmt.Println("index = ", i, "value =", v)
		total += v
	}
	//return 999
	// return total / len(aParam) // ERROR because len(aParam) is NOT of type float64
	return total / float64(len(aParam))
}
