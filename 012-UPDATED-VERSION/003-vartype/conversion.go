package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

func main() {
	a = 45
	b = 46
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	// a = b  // cannot do operations between different types of operands. You need to convert
	a = int(b) // conversion
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
