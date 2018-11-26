package main

import(
  "fmt"
  "bufio"
  "strings"
  "os"
)

func main(){

  const input = "It is not the critic who counts; not the man who points out how the strong man stumbles, or where the doer of deeds could have done them better. The credit belongs to the man who is actually in the arena, whose face is marred by dust and sweat and blood; who strives valiantly; who errs, who comes short again and again, because there is no effort without error and shortcoming; but who does actually strive to do the deeds; who knows great enthusiasms, the great devotions; who spends himself in a worthy cause; who at the best knows in the end the triumph of high achievement, and who at the worst, if he fails, at least fails while daring greatly, so that his place shall never be with those cold and timid souls who neither know victory nor defeat."

  scanner:=bufio.NewScanner(strings.NewReader(input))
  //set the split function for the scannig operation
  scanner.Split(bufio.ScanWords)
  //count the wordlist

  for scanner.Scan(){
    fmt.Println(scanner.Text()) //print each letter
    //in order to learn things, i dont have to learn everyting
    //so just take that above works!!!, i LEGO everything together
  }

 if error:= scanner.Err();error!=nil{
   fmt.Fprintln(os.Stderr,"reading input : ", error)
   //interface appears a lot of times
   /*
   might have to jump to interface to learn before continue
    */
 }
}
