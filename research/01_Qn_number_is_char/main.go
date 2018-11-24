package main

import (
	"fmt"
)

func main() {
	i := 65 //this is ascii of letter A
	/*
		it is A, remember the tutorial about rune
		UTF8,
		all number at just rune
		so if i got a number of decimal, string might be "李"
	*/
	fmt.Println(string(i))

	fmt.Println([]byte("李"))

	a := []byte("李")
	fmt.Printf("%T", a) //[]uint8 // a is a slice of uint8
	fmt.Println()
	/*
		https://stackoverflow.com/questions/11184336/how-to-convert-from-byte-to-int-in-go-programming
			You can use encoding/binary's ByteOrder to do this for 16, 32, 64 bit types
	*/
	//aNumber := binary.BigEndian.Uint64(a) //HOLD
	//fmt.Println("a decimal value = ", aNumber)

	/*
	   var bs []byte
	   value, _ := strconv.ParseInt(string(bs), 10, 64)
	*/
	//aNumberTwo, _ := strconv.ParseInt(string(a), 10, 64)
	//fmt.Println("a decimal value = ", aNumberTwo)

	//b := int(a)
	//google = "golang slice byte to int"

}
