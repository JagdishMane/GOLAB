package main

import (
	"fmt"
)

// Package level vars

var y = 42
var b int

//Declared var z of type string
//Static Programming langauge.  cannot assign string var with int value. var is declared to hold the value of a certain TYPE.
var z = "Shaken, not stirred"
var x string = "Shaken, not x stirred"
var a string = `var using back quite`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println("Printing the value of b", b, "ending")
	fmt.Printf("%T", b)
}
