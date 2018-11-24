package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	slices := [][]byte{
		{1},
		{1, 2},
		{1, 2, 3},
		{1, 2, 3, 4},
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, 5, 6},
		{1, 2, 3, 4, 5, 6, 7},
		{1, 2, 3, 4, 5, 6, 7, 8},
	}

	for _, s := range slices {
		fmt.Println(getInt1(s), getInt2(s))
	}

	a := []byte("李")
	fmt.Printf("%T", a) //[]uint8 // a is a slice of uint8
	fmt.Println()
	fmt.Println("string of a when it is a slice of unit8 = ", string(a))

	aInteger := getInt1(a)
	fmt.Println("integer value of a = ", aInteger) //15113614

	aIntegerTwo := getInt2(a)
	fmt.Println("integer value of a = ", aIntegerTwo) //15113614

	fmt.Println("string of a when it is a integer value", string(aInteger))

	aInteger++
	fmt.Println("string of a when it is a integer value", string(aInteger))
	aInteger++
	fmt.Println("string of a when it is a integer value", string(aInteger))
	// above not displaying "李", as command window cannot show it
	// how to run in a commnad window that display chinese character

	/*
		https://superuser.com/questions/1170656/windows-10-terminal-encoding

			How can I find out the encoding of my windows 10 terminal?
		If by terminal you mean a cmd shell, you can use the chcp command to display and change the code page used.

		Examples

		View the current code page:

		chcp
		Change the code page to Unicode/65001:

		chcp 65001

		65001	utf-8	Unicode (UTF-8)

	*/

	//NOT WORKING
	aStrByteTwo := []byte("李")
	aIntegerThree := getInt1(aStrByteTwo)
	fmt.Println("aIntegerThree = ", aIntegerThree)
	aIntegerThree++
	//convert integer back to slice byte
	//https://www.socketloop.com/tutorials/golang-convert-int-to-byte-array-byte
	fmt.Println(aStrByteTwo) //[230 157 142]

	//aStrByteThree := [230 157 142] //default , parallel type concept in golang

	s := make([]string, 3)
	fmt.Println("emp:", s)
	// We can set and get just like with arrays.

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("emp:", s)

	aStrByteThree := make([]byte, 3)
	fmt.Println("aStrByteThree:", aStrByteThree)
	//aStrByteThree = nil
	aStrByteThree[0] = 230
	aStrByteThree[1] = 157
	aStrByteThree[2] = 142
	fmt.Println("aStrByteThree = ", aStrByteThree)

	fmt.Println("aStrByteThree string = ", string(aStrByteThree))

	aStrByteThree[2] = 143
	fmt.Println("aStrByteThree = ", aStrByteThree)
	fmt.Println("aStrByteThree string = ", string(aStrByteThree))
	aStrByteThree[2]++
	fmt.Println("aStrByteThree = ", aStrByteThree)
	fmt.Println("aStrByteThree string = ", string(aStrByteThree))

	for i := 1; i <= 30; i++ {
		aStrByteThree[2]++
		fmt.Println("aStrByteThree string = ", string(aStrByteThree))
	}
}

func getInt1(s []byte) int {
	var b [8]byte
	copy(b[8-len(s):], s)
	return int(binary.BigEndian.Uint64(b[:]))
}

func getInt2(s []byte) int {
	var res int
	for _, v := range s {
		res <<= 8
		res |= int(v)
	}
	return res
}
