package main

import "fmt"

func main() {
	greeting("lee chuan boon")
	greeting("you can make it")
}

func greeting(aName string) {
	fmt.Println("Greeting", aName)
}
