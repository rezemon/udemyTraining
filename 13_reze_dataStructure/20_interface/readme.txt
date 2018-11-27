

27nov2018 2358pm
https://www.ardanlabs.com/blog/2015/09/composition-with-go.html

Notice the Fasten method requires a value of interface type NailDriver
and we are passing a value of interface type NailDrivePuller.
This is possible because the compiler knows that any concrete type value
that can be stored inside a NailDrivePuller interface value
must MUST MUST MUST also implement the NailDriver interface.
Therefore, the compiler accepts the method call and
the assignment between these two interface type values.

without programming, i am going to try to interpret this according to my own
type Printer interface {
  Print()
}
to be a printer must implement method Print

type Fax interface {
  Fax()
}

PrinterFax interface {
  Printer
  Fax
}

if we have a
type Struct Employee{
}
and Employee has a func
//func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board) {
func (Employee) SendFax(f Fax){

}
the param take a param of Fax interface
if we passed in a PrinterFax interface, it will still work
This is possible because the compiler knows that any concrete type value
that can be stored inside a PrinterFax interface value
must MUST MUST MUST also implement the Fax interface.
Therefore, the compiler accepts the method call and
the assignment between these two interface type values.

===

what about the concept
ANY type is a interface
because interface is a type
int is a type
string is a type
if interface is declared as interface{}
e.g.
type interface{

}
or rather
var AAA interface{
   // NO METHODS DEFINED
   // NO METHODS DEFINED
   // NO METHODS DEFINED
} , AAA is a empty interface
and because it implement NO methods
thus
int type is also a empty interface{} because it implement NO methods, hence it is a empty interface
string type is also a empty interface{} because it implement NO methods, hence it is a empty interface
etc etc

so if int is a type, it is also a interface
if
type MyInt struct{
    value int
}
if this "MyInt" implements all the methods found in "Int type"
we can say that "MyInt" is a integer interface
Ducking, if it quack like a duck, it is a duck
so if Duck interface required 100 methods
if a Struct implement ALL 100 methods, it is a Duck
if a Struct implement > 100 methods, it is also a Duck
if a Struct implement 99.99 methods, it is NOT NOT NOT a Duck
and this is the concept of Duck typing

==
composition
today 27nov2018, talked to tian hoe, learn something abt interface ..
hmm..should be in interface
composition

type Printer interface {
 methodA
 methodB
 ...
 methodZ
}
if something that implement the methods declarded in Printer interface , then that "something" is a Printer

likewise
type Fax interface{
 method1
 method2
 ...
 method999999
}
if something that implement the methods declared in Fax interface, then that "something" is a Fax

composition then become clear
if it is a PrinterFax machine ..we need a PrinterFax interface
PrinterFax interface{
     Printer
     Fax
}
################################
the beauty is in its simplicity
################################
==
understanding interface

if it quack like a duck, it is a duck
means i have implemented a duck
this is a statement to remember for interface
