package main

import(
  "fmt"
  "net/http"
  "bufio"
"os"
  "log"
  //"io/ioutil"

)

//get all the letter in english language and print them out
func main(){
  //http://www.angelfire.com/extreme4/safer_sephiroth/EVERY_WORD_EVER.htm
  //http://www.angelfire.com/extreme4/safer_sephiroth/EVERY_WORD_EVER.htm
//http://www.mhhe.com/mayfieldpub/maner/resources/actb.txt
//http://www.mhhe.com/mayfieldpub/maner/resources/actb.txt

  //response,error := http.Get("http://www.mhhe.com/mayfieldpub/maner/resources/actb.txt")
  //http://www-personal.umich.edu/~jlawler/wordlist
  response,error := http.Get("http://www-personal.umich.edu/~jlawler/wordlist")

  if error != nil {
    log.Fatalln(error)
  }

/*
  bs,_ := ioutil.ReadAll(response.Body)
  str := string(bs)
  fmt.Println(str)
  */
 //words :=make(map[string]string,0)
 words :=make(map[string]string)

 sc := bufio.NewScanner(response.Body)
sc.Split(bufio.ScanWords)

for sc.Scan(){
  if sc.Text() == "happy"{
    words[sc.Text()] = "is the motto of lee chuan boon"
  } else {
    words[sc.Text()] = ""
  }

}

if error:= sc.Err(); error!= nil {
  fmt.Fprintln(os.Stderr,"reading input", error)
}

i := 0
for k, _ := range words{
  fmt.Println(k)
  if i == 20{
    break
  }
  i++
}

fmt.Println(words["apple"])
value,okay := words["happy"]
fmt.Println("happy : " ,value , " - " , okay)

value2,okay2 := words["happy123"]
fmt.Println("happy123 : ", value2 , " - " , okay2)

}
