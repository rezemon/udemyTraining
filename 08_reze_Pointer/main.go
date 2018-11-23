package main

import "fmt"

func main() {

	a := "lee chuan boon"

	fmt.Print("pointer value of a = ")
	fmt.Printf("%d\n", &a)
	fmt.Print("a is of type =")
	fmt.Printf("%T\n", &a)
	//above will show *string ==> it is a pointer to type "string"

	// declare something that is of type pointer
	// *string ==> this is a pointer to an string

	var b *string
	b = &a
	fmt.Print("pointer value of b (HEX) = ")
	fmt.Print(b)
	fmt.Println()
	fmt.Print("pointer value of b (DEC) = ")
	fmt.Printf("%d", b)
	fmt.Println()
	fmt.Print("b is of type =")
	fmt.Printf("%T", b)
	/*
		   b is REFERENCING to address &a
		   it is a pointer to a string
		   how to DE-REFERENCE it??
		   use *b
			 THAT IS EQUIVALENT TO SAYING, dont give me the pointer address
			 give me the value of that pointer address
	*/
	fmt.Println()
	fmt.Println("value of b = ", *b)

	/*
		now, change the value of THE ADDRESS
	*/
	*b = "lee chuan boon is VERY VERY VERY HAPPY"
	fmt.Println()
	fmt.Println("value of a is changed to : ", a)

	/*
		EVERYTHING IN GOLANG IS PASSED BY value
		by value means PASS THE MEMORY address
		this is different from what i learn, PASS BY REFERENCE (pointer), PASS BY VALUE (make another copy)
		so, here TWEET a bit, PASSED BY VALUE IN GO <== EQUIVALENT ==> PASSED BY REFERENCE,
		ABOVE WRONG WRONG WRONG XXXXXXXXX

		everything in GOLANG is PASSED BY value
		even if we passed a memory address, it is passing by VALUE
		we PASS A VALUE, GOLANG discourage to say it is making a copy
		we pass "lee chuan boon", it is pass by VALUE
		we pass a string pointer to "lee chuan boon", the "string pointer" is PASSED BY VALUE, GOT IT GOT II !!!!!!

	*/

	//above if a:=1234
	/*
		then type of &a will be "*int"
		which is a pointer to type "int"
	*/
}
