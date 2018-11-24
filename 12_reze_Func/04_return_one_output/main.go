package main

import "fmt"

func main() {
	fmt.Println(greeting("lee", "chuan boon"))
	fmt.Println(greeting2("lee", "chuan boon")) //dont have space, golang docu says have space added
	fmt.Println(greeting3("lee", "chuan boon"))

}

func greeting(surname string, name string) string {
	return (name + " " + surname)
}

func greeting2(surname string, name string) string {
	return fmt.Sprint(name, surname)
}

func greeting3(surname string, name string) string {
	//return (name, surname) //THIS DONT WORK
	return "dont work"
}
