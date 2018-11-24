package main

import "fmt"

func main() {
	greeting("lee", "chuan boon")
	greeting("tan", "kim choo")
}

func greeting(surname string, name string) {
	// US notation of greeting
	fmt.Println("Hello", name, surname)
}
