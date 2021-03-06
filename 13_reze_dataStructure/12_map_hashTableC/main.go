package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	//"os"
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
//	response, error := http.Get("http://www-personal.umich.edu/~jlawler/wordlist")

//http://25.io/toau/audio/sample.txt

//response, error := http.Get("https://ebooks.adelaide.edu.au/m/melville/herman/m53m/chapter10.html")
//https://ocw.mit.edu/ans7870/6/6.006/s08/lecturenotes/files/t8.shakespeare.txt
//https://ocw.mit.edu/ans7870/6/6.006/s08/lecturenotes/files/t8.shakespeare.txt
//response, error := http.Get("http://25.io/toau/audio/sample.txt")
response, error := http.Get("https://ocw.mit.edu/ans7870/6/6.006/s08/lecturenotes/files/t8.shakespeare.txt")
	if error != nil {
		log.Fatalln(error)
	}

//scan the page
scanner := bufio.NewScanner(response.Body)
defer response.Body.Close()

//set the split function for the scanning operation
scanner.Split(bufio.ScanWords)

//create slice to hold counts
buckets := make([]int,200)
//loop over the words
for scanner.Scan(){
	n:= HashBucket(scanner.Text())
	buckets[n]++
}

/*
as can see, for hashing, we need a even distribution to be EFFICIENT

 */
fmt.Println(buckets[65:123])

fmt.Println("##############################")
for i:=28; i <=126; i++{
	fmt.Printf("%v - %c - %v \n",i,i,buckets[i])
}
}

//func HashBucket(word string)
func HashBucket(word string) int{
	return int(word[0])
}
