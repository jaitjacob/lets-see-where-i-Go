package main

import "fmt"

func main() {
	fmt.Println("This is a String value")
	fmt.Println("This is an integer concatenated to a String: ", 9)
	fmt.Println("This is a Boolean concatenated to a String: ", true)

	//Boolean operators
	fmt.Println(true && false) //AND
	fmt.Println(true || false) //OR
	fmt.Println(!true)         //NOT

	//Easter Egg: The MCP ended most of its conversations with Dillinger and Sark with the phrase "End of Line."
	fmt.Println("End of Line.")
}
