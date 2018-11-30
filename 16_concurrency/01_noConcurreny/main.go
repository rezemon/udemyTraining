package main

import (
	"fmt"
)

func main() {
	foo()
	bar()
}
func foo() {
	for i := 0; i <= 30; i++ {
		fmt.Println(i, "foo")
	}
}

func bar() {
	for i := 0; i <= 30; i++ {
		fmt.Println(i, "bar")
	}
}
