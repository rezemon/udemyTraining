database STRUCT

it is so intuitive
i dont have to look at it twice

golang spec
effective go, they dont talk about struct ?? dont know why

composition => this is new
composing program
composition is something you do with golang

Embedding related to Struct
need to learn how effective go talk about it

Struct can have anonymouse field

in STRUCT
a field declared with a type but NO EXPLICIT field name is an anonymouse field.
it is also called an embedded field
OR
an embedding of the type in the STRUCT
an embedded type must be specificed as a type surname
OR
as a pointer to a non-interface type name *T
and T itself may not be a pointer TYPE
The unqualified type name acts as the field name

overriding is called "promotion" in golang

golang is Object oriented

###########################
1. encapsulation in golang
###########################
we create a STRUCT can called it a type
we declared a variable of that type  (instantiation)
behaviour is implemented by the fact that we can attached a func to a type
and that can be named as METHOD
we declared STRUCT, when we need to store the state of an object
in golang
private ==> NOT EXPORTED
public ==> EXPORTED
and we should avoid saying, private or public

###############################
2. reusability or INHERITANCE
###############################
type Person {
  Name string
  Age int
}

type SecretAgent{
  Person (reze: this is embedded type, no field name, but have type)
  (the inner type Person got promoted to the outer type SecretAgent)
  LicenseToKill bool
}

above, SecretAgent inherit from Person
therefore SecretAgent can ACCESS all the fields in Person type
and this is INHERITANCE

##############################
3. Polymorphism
##############################
we achieve this via interface
(reze : i have to read up a lot on interface )
e.g. reader interface, ..

due to too many mentioning of interface, i need to go read this up
and have a SUPER understanding of it because i continue

==
composition
today 27nov2018, talked to tian hoe, learn something abt interface ..
hmm..should be in interface
