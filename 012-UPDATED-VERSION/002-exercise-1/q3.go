package main

import "fmt"

var x int = 43
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%s\t%t\t", x, y, z)
	fmt.Println(s)
}
