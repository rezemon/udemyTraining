package main

import(
  "fmt"
  "log"
  "io/ioutil"
  "net/http"
)

func main(){

  response, error := http.Get("https://ocw.mit.edu/ans7870/6/6.006/s08/lecturenotes/files/t8.shakespeare.txt")
  if error != nil{
    log.Fatal(error)
  }

  page,error := ioutil.ReadAll(response.Body) //page could be byte slice
  response.Body.Close()

  if error != nil{
    log.Fatal(error)
  }

  fmt.Printf("%s",page) //take the byte slice and print it as a string
  
}
