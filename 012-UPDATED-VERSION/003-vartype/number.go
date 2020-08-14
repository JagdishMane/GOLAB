package main

import "fmt"

var a int = 19
var b float64 = 10.2082034808

func main() {
	x := 24
	//x = 24.03 //error as x is type int
	y := 04.199898
	fmt.Println(x, "\n", y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(a, "\n", b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
}
