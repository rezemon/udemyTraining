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
	//response, error := http.Get("http://www-personal.umich.edu/~jlawler/wordlist")
	response, error := http.Get("https://ocw.mit.edu/ans7870/6/6.006/s08/lecturenotes/files/t8.shakespeare.txt")

	if error != nil {
		log.Fatal(error)
	}

	//scan the page
	//NewScanner takes a reader and response.Body implements the reader interface
	// scan the page
	// NewScanner takes a reader and res.Body implements the reader interface (so it is a reader)
	/*
		i have to understand interface, 26nov2018 2221pm
		i just talk to tianhoe this afternoon
		what is a interface, think in terms of printer
		in microsoft OS,
		for printing, microsoft define a print function called interface..
		it requires all vendor who want to put their printer to be used on MS win10, to implement this print interface
		otherwise it wont work in microsoft win10
		and this is logical, cause MS might need to deal with FEW thousands vendor, and it cannot handle all special hardware
		THEREFORE
		in order for a printer manufacturer to print in MS win10,
		they have to implement a print interface ..
		so to align with the author or teacher
		if NewPrinter takes a printer and a print buffer (which contains data to be printed) to implement the print interface , IT IS A printer
		IF IT QUACKS !!!, IT IS A DUCK
		IF IT FLY !!!, IT IS A PLANE
		IF IT FLY !!!, IT IS A BIRD
		IF IT FLY !!!, IT IS SUPERMAIN
		IF IT SWIM, IT IS A FISH
	*/
	scanner := bufio.NewScanner(response.Body)
	defer response.Body.Close()

	//set the split function for the scanning, WHAT IS THIS ???
	scanner.Split(bufio.ScanWords)

	//loop over the words
	for scanner.Scan() {
		fmt.Println("scan text =", scanner.Text())
	}

	if error := scanner.Err(); error != nil {
		fmt.Fprintln(os.Stderr, "reading input : ", error)

	}

}
