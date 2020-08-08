package main

import "fmt"
//x = 22
//y == 32.2
//uint8 0 to 255
//int8 -128 to 127
/int64
var z int8 = -125
var u uint8 = 255
//var z int

func main() {
	x := 22
	y := 32.2
	fmt.Println(x, y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
}
