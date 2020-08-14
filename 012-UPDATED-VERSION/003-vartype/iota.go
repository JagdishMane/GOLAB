package main

import (
	"fmt"
)

const (
	a = iota
	b
	c
)
//const reset the iota. iota fills the successive values.
const (
	d = iota
	e
	f
)

func main() {

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}