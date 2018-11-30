Once a type implement an Interface.
the type in this example is "EpsonPrinter"
an ENTIRE world of functionality can be opened up
to the values of that type "EpsonPrinter"

type Printer interface {
  Printing() string
}

type struct EpsonPrinter {
 buffer string
}

func (e EpsonPrinter) Printing() string{
  // some implementation here
  // some codes here
}

interfaces are type that just declare behaviour,
this behaviour is NEVER implemented by the interface type DIRECTLY
but instead
it is implemented by user-defined types VIA methods
when a user-defined type implements the set of methods declared by an interface type ,
values of the user-defined type can be assigned to the value of the interface type
ABOVE MEANS
aEpsonPrinter := EpsonPrinter{"some buffer string"}
var aPrinterInterface Printer //define printer interface

value of the interface type (aPrinterInterface)
value of user-defined type (aEpsonPrinter)
aPrinterInterface = aEpsonPrinter
value of user-defined type ASSIGNED to value of interface type

if a method call is make against an interface value (Printer)
the EQUIVALENT method (in this case Printing() defined in the func above ) for the stored user-defined value is executed

Since ANY user-defined type can implement ANY interface
methods calls against an interface are POLYMORPHIC

e.g.
type Shape Interface{
  area() string
}

struct Square{

}
struct Circle{

}

//Square implement Shape interface
func (s Square) area() string{
  return area of the square
}

//Circle implement Shape interface
func (c Circle) area() string{
  return area of the circle
}

Square call against Shape interface execute Square area()
Circle call against Shape interface execute Circle area()
