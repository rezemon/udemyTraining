package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	myResult, myError := http.Get("http://www.example.com")
	//myResult, myError := http.Get("http://www.mcleods.com")
	//myResult, myError := http.Get("http://www.stoneforest.com.sg/")
	if myError != nil {
		log.Fatal(myError)
		fmt.Println("Error")
		fmt.Printf("%s\n", myError)
	}

	myPage, myError := ioutil.ReadAll(myResult.Body)
	myResult.Body.Close()
	if myError != nil {
		log.Fatal(myError)
		fmt.Println("Error")
		fmt.Printf("%s\n", myError)
	}

	fmt.Printf("%s\n", myPage)

	fmt.Printf("%T\n", myResult)
	fmt.Printf("%T\n", myPage)
}
