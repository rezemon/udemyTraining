package main

import(
  "fmt"
)

func main(){
  hashBucketNumber := HashBucket("Go",12)
  fmt.Println(hashBucketNumber)

  hashBucketNumberB := HashBucketB("Go",12)
  fmt.Println(hashBucketNumberB)
}

func HashBucket(word string,numberOfBucket int32) int32 {
  var x int32 //alias for rune
  x = rune(word[0])
  return x % numberOfBucket

}

func HashBucketB(word string, numberOfBucket int) int{

  x := int(word[0]) // take a word[0] ==> a rune, convert to an integer int(word[0])
  return x % 12
}
