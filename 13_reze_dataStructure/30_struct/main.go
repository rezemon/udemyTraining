//golang STRUCT
//encode into a JSON stream

//ENCODE
//we want to send something out
//we going to write it to somewhere
//and USUALLY we use the writer interface to write

//DECODE
//something is coming in
//we going to read it
//and USUALLY we use the reader interface to read

//marshal unmarshal normally stay within the program

//ENCODE and DECODE
//something is coming in, something is going out

//if something is going out, i am going to encode it, and i am going to write it
//if something is coming in, i am going to read it, i am going to decode it,

package main
import(
  "encoding/json"
  "fmt"
  "os"
)

type Person struct{
  Surname string
  Name string
  Age int
}

type PersonB struct{
  Surname string
  Name string
  Age int
}

func main()  {
  fmt.Println("ENCODING")
  p1:= Person{"lee","chuan boon",33}
  //json.NewEncoder(w io.Writer) ??

  //json.NewEncoder(os.Stdout).Encode(p1)

  /*
  // NewEncoder returns a new encoder that writes to w.
  func NewEncoder(w io.Writer) *Encoder {
  	return &Encoder{w: w, escapeHTML: true}
  }
#################################################
REZE above return a new *Encoder that writes to w
#################################################
*/
aNewEncoder := json.NewEncoder(os.Stdout)
// NOW NOW NOW we have a POINTER to the aNewEncoder
//we writing it to standard out

/*
#################################################
The new *Encoder has method Encode that take "v interface{}"
"v interface{}" is normally a golang STRUCT
#################################################
  func (enc *Encoder) Encode(v interface{}) error {
  	if enc.err != nil {
  		return enc.err
  	}
  	e := newEncodeState()
  	err := e.marshal(v, encOpts{escapeHTML: enc.escapeHTML})
  	if err != nil {
  		return err
  	}
   */
aNewEncoder.Encode(p1)

//above also can be written in the following manner
p2 := Person{"wong","ken meng",34}
json.NewEncoder(os.Stdout).Encode(p2)
//above , execute, read from LEFT to RIGHT

//so now what is a writer
//this will write to a place , so we can send it out

}
