package main

import (
	"fmt"
)

func main() {
	fmt.Println([]byte("Hello"))
	fmt.Println([]byte("世界"))  //[228 184 150 231 149 140]
	fmt.Println([]byte("李全文")) //[230 157 142 229 133 168 230 150 135]

	//reverse
	//李 is presented by [230 157 142] 3 bytes
	//my first try

	//aStr := "230 157 142"
	//aStrByte := []byte(aStr)

	//fmt.Println(aStrByte)
	aStrByte := []byte("李全文")

	fmt.Print("aStrByte value = ")
	fmt.Println(aStrByte)

	fmt.Print("aStrByte is of type = ")
	fmt.Printf("%T", aStrByte)

	fmt.Println()
	fmt.Print("the string representation = ")
	fmt.Println(string(aStrByte))

	//with this , i can print all the unicode for chinese character
}
