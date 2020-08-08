package main

import (
	"fmt"
)

var y = 43
//Declare these is a variable with identifies z  is of type z and assigned values 0 to z.
var z int 
//Each element of such a variable or value is set to the zero value for its type: false for booleans, 0 for numeric types, 
//"" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(y)
	foo()
	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)
}