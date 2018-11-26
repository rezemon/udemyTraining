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
/*
 create slice of the slice of string to hold slices of words
 [][]string 12
 first "[]", we are creating slice of size =12
 the 12 slice will contain []string
 EQUIVALENT TO SAYING
 the 12 slice will contain slice of type string
 so EQUI to what the tutorial teach
 create slice to CONTAIN "slice of string"
 slice of string will hold the words (slices of word)
*/
buckets := make([][]string,12) // make is ok, it wont be nil
//for slice, even if it is nil, it is OK, we can use APPEND to add things
//but for MAP, cannot be a nil, as there is no function to add to a "map type"
// create the slices to hold the words
for i:=0;i<=12;i++{
	buckets = append(buckets,[]string{}) // initializing
}
	//loop over the words
	for scanner.Scan() {
		//fmt.Println("scan text =", scanner.Text())
		word := scanner.Text()
		n:= HashBucketNumber(word, 12) //the first slice []
		//divided into 12 buckets
		//the first slice has index 0 to index 11
		//based on the hashBuckNumber return
		//at the first [], we use [n] to access to the bucket
		//then put the words there
		buckets[n] = append(buckets[n], word)
	}

//print the len of each buckets
for i:=0;i<=12;i++{
	fmt.Println("length of bucket [", i, "] =" , len(buckets[i]))
}

//print the words in one of the buckets
if false{
fmt.Println(buckets[6]) //too many things //print a slice of it
}
//fmt.Println(buckets[3])
//fmt.Println(buckets[1])

aSlice := buckets[6][20:30] //take the 20th element up to and BEFORE 30th element
// 20th 21th ...29th
for index,value := range(aSlice){
	 fmt.Println(index , " - " , value)
}

	if error := scanner.Err(); error != nil {
		fmt.Fprintln(os.Stderr, "reading input : ", error)

	}

}

func HashBucketNumber(word string, numOfBuckets int) int {

  sum := 0
	// range over word
	// word is make up of RUNE
	//word is make up of RUNE
	//word is make up of RUNE
	//so when we range over it, we get rune
	for _,value := range word{
		sum += int(value) // "A" ==> 65, "B" ==>66...
	}
	return sum% numOfBuckets
}
