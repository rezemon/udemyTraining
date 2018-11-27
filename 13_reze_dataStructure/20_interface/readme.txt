
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
