package main

import "fmt"

func main() {

	var aMap = make(map[string]string, 0)

	aMap["zero"] = "good morning"
	aMap["one"] = "good afternoon"
	aMap["two"] = "good evening"
	aMap["three"] = "good night"

	fmt.Println("aMap = ", aMap)

	//to range loop thru a map to list all elements

	for key, value := range aMap {
		fmt.Println("key = ", key, " - ", "value = ", value)
	}

}
