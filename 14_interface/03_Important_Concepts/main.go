package main

import(
  "fmt"
)

type Person struct{
  //empty struct
}

type PersonWithName struct{
  Name string
}
type PersonPtr *struct{

}
//a empty interface{}, means can take any value
type I interface{
  // I is empty interface
  // no methods
  // thus all TYPE are empty interface
}

func SayHello() string{
  return "hello world, here i come, Lee Chuan Boon"
}

//how to point to a func
func SayHelloB() string{
  return "hello yes yes yes"
}

//i cannot point to a func
//can i point to a method of a STRUCT

type HPLaserJet1000 struct{
  pages string
}

type Brother8840D struct{
  pages string
}

type Canon123 struct{
  pages string
}

type Printer interface{
  Printing(aString string) string
}

type Fax interface{
  Faxing(aString string) string
}

//because HPLaserJet1000 struct implement Printing method, therefore HPLaserJet1000 is a Printer
func (hplj HPLaserJet1000) Printing(aStr string) string{
    hplj.pages = aStr
   return aStr
}

//because Brother8840D struct DID NOT implement Printing method, therefore Brother8840D is NOT a Printer
func (b  Brother8840D) xPrinting(aStr string) string {
  b.pages = aStr
  return aStr
}

// Canon123 is a Pointer to the Canon123 Struct
func (c *Canon123) Printing(aStr string) string{
  c.pages = aStr
  return aStr
}

func main()  {
  //https://godoc.org/golang.org/x/tools
  //https://godoc.org/golang.org/x/tools
  fmt.Println("how to use describe")
  var aInt int
  aInt = 33
  // a int is also a empty interface
  describe(aInt)

  aStr := "hello"
  aFloat := 1.234
  aBool := true
  aStruct := Person{}
  aStructName := PersonWithName{"lee chuan boon"}
  var aStructB PersonPtr   //? dont understand yet

  //we can put a slice of interface to hold all this
  //but now we try to just call describe on it own
  describe(aStr)
  describe(aFloat)
  describe(aBool)
  describe(aStruct)
  describe(aStructB)

fmt.Println("===========================")
  //define a slice of interface
  //a slice of interface
  aSliceInterface := make([]interface{},20)
  aSliceInterface = append(aSliceInterface,aInt)
  aSliceInterface = append(aSliceInterface,aStr)
  aSliceInterface = append(aSliceInterface,aFloat)
  aSliceInterface = append(aSliceInterface,aBool)
  aSliceInterface = append(aSliceInterface,aStruct)
  aSliceInterface = append(aSliceInterface, aStructName)
  aSliceInterface = append(aSliceInterface, aStructB)

  var aIntPtr *int // this declare aIntPtr is of type integer pointer
// it is pointing to a integer value
  aSliceInterface = append(aSliceInterface, aIntPtr)
  aTempInt := 3388 //value
  aIntPtr = &aTempInt //point the address of 3388 to the integer pointer
  aSliceInterface = append(aSliceInterface, aIntPtr)
  aSliceInterface = append(aSliceInterface, *aIntPtr)
  fmt.Println("value of aTempInt", *aIntPtr)

  var aStrPtr *string //string is a type
  aTempStr := "lee chuan boon"
  aStrPtr = &aTempStr
  aSliceInterface = append(aSliceInterface, aStrPtr)
  aSliceInterface = append(aSliceInterface, *aStrPtr)

  //string is a type
  //Person is also a type, it is a Struct type or so called User Defined Type
  var aStructPtr *PersonWithName
  aTempStructPtr := PersonWithName{"wong ken meng"}
  aStructPtr = &aTempStructPtr
  aSliceInterface = append(aSliceInterface, aStructPtr)
  aSliceInterface = append(aSliceInterface, *aStructPtr)

/*
to point to something, it has to be a TYPE ???
 */
 //var aFuncPtr *SayHello // yes it can point to a function //ERROR SayHello is not a type
 // var aFuncPtr *SayHello() // ERROR
 //aSliceInterface = append(aSliceInterface,aFuncPtr) //it point to the function SayHello
 //it store the address that point to the function SayHello

var aInterface I //above I is a empty interface
// now i pass a interface into itself
aSliceInterface = append(aSliceInterface,aInterface) //this not printed !!! because it is nil , empty interface is nil ???

fmt.Println("Empty interface type BEGIN")
fmt.Printf("%T\n", aInterface)
fmt.Println("Empty interface type END")

var aPrinterInterface Printer
aSliceInterface = append(aSliceInterface,aPrinterInterface) //this also nil


var aPrinterInterfaceB Printer
aHpPrinter := HPLaserJet1000{"print i am a HP laser jet 1000"}
aPrinterInterfaceB = &aHpPrinter // WORK!!
aPrinterInterfaceB = aHpPrinter // also WORK!!
tempStr := ""
tempStr = aPrinterInterfaceB.Printing("i am printer interface")
fmt.Println(tempStr)
fmt.Println("HP Printer interface type BEGIN")
fmt.Printf("%T\n", aPrinterInterfaceB)
fmt.Println("HP Printer interface type END")

aSliceInterface = append(aSliceInterface, aPrinterInterfaceB)
/*
# github.com/GoesToEleven/udemyTraining/14_interface/03
.\main.go:151:20: cannot use aBrotherPrinter (type Brother8840D) as type Printer in assignment:
	Brother8840D does not implement Printer (missing Printing method)
 */
//aBrotherPrinter := Brother8840D{"printing i am a Brother MFC 8840D"}
//aPrinterInterfaceB = aBrotherPrinter
//REZE above dont work, because Brother8840D DOES NOT IMPLEMENT Printing method

/*
# github.com/GoesToEleven/udemyTraining/14_interface/03
.\main.go:173:20: cannot use aCanonPrinter (type Canon123) as type Printer in assignment:
	Canon123 does not implement Printer (Printing method has pointer receiver)
 */
//aCanonPrinter := Canon123{"print i am a Canon 123 printer"}
//aPrinterInterfaceB = aCanonPrinter
//REZE above dont work, because Canon123 DOES NOT IMPLEMENT Printing method
//WHY ??
//Printing method has pointer receiver
//instead of giving it the address of the STRUCT , i give the value
//the compiler not able to determine
//like in the Brother8840D. it search through and know that Brother8840D did not implement Printing method
//in the Canon123, it CANNOT CANNOT CANNOT POINT to the Canon123 struct WITHOUT the address
//hence compiler says Canon123 DOES NOT IMPLEMENT Printing method
//solution as below
//
aCanonPrinter := Canon123{"print i am a Canon 123 printer"}
aPrinterInterfaceB = &aCanonPrinter //pass address, so that Compiler can find the Canon123 struct
fmt.Println("Canon Printer interface type BEGIN")
fmt.Printf("%T\n", aPrinterInterfaceB)
fmt.Println("Canon Printer interface type END")
aSliceInterface = append(aSliceInterface, aPrinterInterfaceB)




  for index,value := range aSliceInterface{
    if value != nil {
    describe(aSliceInterface[index])
    }
}
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)

}
