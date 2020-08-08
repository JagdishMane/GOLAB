package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)   // type
	fmt.Printf("%b\n", y)	// binary
	fmt.Printf("%x\n", y)	// Hex
	fmt.Printf("%#x\n", y)  // add ox
	//fmt.Printf("%T\t%b\t%x\t",y ,y, y)
	s := fmt.Sprintf("%T\t%b\t%x\t",y ,y, y)
	fmt.Println(s)
	fmt.Printf("%v", y)
	}