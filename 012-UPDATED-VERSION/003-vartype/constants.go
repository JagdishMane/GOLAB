package main

import (
	"fmt"
)

const (
	a = 33
	b = 43
	c = "Jagdish"
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	a := 44  // if you put a = 44 it fails as a is constant
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}