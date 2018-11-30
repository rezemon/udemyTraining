package main

import(
  "fmt"
  "net/http"
  "log"
  "io/ioutil"
)

//get all the letter in english language and print them out
func main(){
  //http://www.angelfire.com/extreme4/safer_sephiroth/EVERY_WORD_EVER.htm
  //http://www.angelfire.com/extreme4/safer_sephiroth/EVERY_WORD_EVER.htm
  //http://www.mhhe.com/mayfieldpub/maner/resources/actb.txt
  //http://www.mhhe.com/mayfieldpub/maner/resources/actb.txt

  // response,error := http.Get("http://www.angelfire.com/extreme4/safer_sephiroth/EVERY_WORD_EVER.htm")
  //https://sites.google.com/a/vhhscougars.org/johnsearch/searchindex/oxford-english-dictionary/Oxford_English_Dictionary.txt
  //https://sites.google.com/a/vhhscougars.org/johnsearch/searchindex/oxford-english-dictionary/Oxford_English_Dictionary.txt
  //found

  //response,error := http.Get("http://www.mhhe.com/mayfieldpub/maner/resources/actb.txt")
//http://www-personal.umich.edu/~jlawler/wordlist
//http://www-personal.umich.edu/~jlawler/wordlist

//  response,error := http.Get("https://sites.google.com/a/vhhscougars.org/johnsearch/searchindex/oxford-english-dictionary/Oxford_English_Dictionary.txt")
 response,error := http.Get("http://www-personal.umich.edu/~jlawler/wordlist")
  if error != nil {
    log.Fatalln(error)
  }

  bs,_ := ioutil.ReadAll(response.Body)
  str := string(bs)
  fmt.Println(str)
}
