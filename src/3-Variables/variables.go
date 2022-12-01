package main

import "fmt"

func main() {
	// 'var' declares one or more variables
	var name = "jxsnow"
	fmt.Println(name)

	// ':=' is the short variable declaration assignment
	pi := 3.14
	fmt.Println(pi)

	//There are multiple way to declare a variable in Go.

	//Way 1: Declare a variable called i of data type int without initialization.
	var i int //Only use long form, var i int, when you’re not initializing the variable.
	fmt.Println("Declare a variable called i of data type int without initialization.")
	fmt.Println(i) // print zero because integer type variable i is a non value(not initialized with a value)

	//Following prints Zero Values for int, string, float and bool.
	var l int
	var m string
	var n float64
	var o bool

	fmt.Printf("var l %T = %+v\n", l, l)
	fmt.Printf("var m %T = %q\n", m, m)
	fmt.Printf("var n %T = %+v\n", n, n)
	fmt.Printf("var o %T = %+v\n\n", o, o)

}
