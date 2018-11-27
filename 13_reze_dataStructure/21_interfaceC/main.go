//

/* ORINGAL
package main

import (
	"fmt"
)

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func main() {
	names := []string{"stanley", "david", "oscar"}
	PrintAll(names)
}
prog.go:15:10: cannot use names (type []string) as type []interface {} in argument to PrintAll
Go build failed.
if change to below , it is ok
i change it myself
source : http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
 */
package main

import (
	"fmt"
)

//func PrintAll(vals []interface{}) {
func PrintAll(vals interface{}) {
	//for _, val := range vals {
	//	fmt.Println(val)

	//}
	fmt.Println(vals)
	fmt.Printf("%T",vals)
}

func main() {
	//names := []string{"stanley", "david", "oscar"}
	//PrintAll(names)
	PrintAll("lee")
	/*
	what happen is "lee" is passed in as a string
	string is a "type"
	and string is able to implement empty interface{}
	so empty interface{} assume the type of "string"
	go printf("%T",xxx) and see what is the type of xxx
	 */
}
