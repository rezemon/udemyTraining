package main

import "fmt"

func main() {

	a := 33
	fmt.Println("a =", a)
	fmt.Println("address of a = ", &a)

	//b := 66
	fmt.Printf("%d\n", &a)

	var aString string
	fmt.Print("please input your name : ")
	fmt.Scan(&aString)
	fmt.Println("Hello ", aString, "Welcome to GoLang")
	fmt.Print("please input your name : ")
	fmt.Scanln(&aString)
	fmt.Println("Hello ", aString, "Welcome to GoLang AGAIN")
}
