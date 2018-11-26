package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	//"io/ioutil"
)

//get all the letter in english language and print them out
func main() {
	//http://www.angelfire.com/extreme4/safer_sephiroth/EVERY_WORD_EVER.htm
	//http://www.angelfire.com/extreme4/safer_sephiroth/EVERY_WORD_EVER.htm
	//http://www.mhhe.com/mayfieldpub/maner/resources/actb.txt
	//http://www.mhhe.com/mayfieldpub/maner/resources/actb.txt

	//response,error := http.Get("http://www.mhhe.com/mayfieldpub/maner/resources/actb.txt")
	//http://www-personal.umich.edu/~jlawler/wordlist
	response, error := http.Get("http://www-personal.umich.edu/~jlawler/wordlist")

	if error != nil {
		log.Fatalln(error)
	}

	/*
	   bs,_ := ioutil.ReadAll(response.Body)
	   str := string(bs)
	   fmt.Println(str)
	*/
	//words :=make(map[string]string,0)
	words := make(map[string]string)

	sc := bufio.NewScanner(response.Body)
	sc.Split(bufio.ScanWords)

	tempCounter := 0
	for sc.Scan() {

		if tempCounter <= 3 {
			fmt.Println("scan text =", sc.Text())
		}

//this before the switch as it has a else case that set everything to ""
		if sc.Text() == "happy" {
			words[sc.Text()] = "is the motto of lee chuan boon"
		} else {
			words[sc.Text()] = ""
		}


		switch sc.Text() {
		case "success":
words[sc.Text()] = "will be soon"
			fmt.Println("found success = " , sc.Text())

		case "receptive":
			fmt.Println("found receptive = ", sc.Text())
			words[sc.Text()] = "open minded to learn new things"

case "car":
	words[sc.Text()] = "jordon says has remote control"
		}

		tempCounter++
	}

	if error := sc.Err(); error != nil {
		fmt.Fprintln(os.Stderr, "reading input", error)
	}

	i := 0
	for k, _ := range words {
		fmt.Println(k)
		if i == 20 {
			break
		}
		i++
	}

	fmt.Println(words["apple"])
	value, okay := words["happy"]
	fmt.Println("happy : ", value, " - ", okay)

	value2, okay2 := words["happy123"]
	fmt.Println("happy123 : ", value2, " - ", okay2)

	value3, okay3 := words["success"]
	fmt.Println("success : ", value3, " - ", okay3)
	value4, okay4 := words["receptive"]
	fmt.Println("receptive : ", value4, " - ", okay4)

	value5, okay5 := words["car"]
	fmt.Println("car : ", value5, " - ", okay5)
}
