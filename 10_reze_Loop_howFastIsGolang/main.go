package main

import "fmt"

func main() {
	x := 1
	for {
		fmt.Println(x)
		x++
		if x > 1000 {
			break
		}
	}
}
